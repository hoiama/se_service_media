package sei_video_webrtc

import (
	"fmt"
	"github.com/pion/rtcp"
	"io"
	"log"
	"sei_service_video/main/src/sei_helper/errorHelper"
	"time"
)


// admin is the caller of the method
// if admin connects before peer: since admin is first, admin will create the channel and track and will pass the track to the channel
// if user connects before admin: since user came already, he created the channel and is listening and waiting for admin to create and pass track
func createTrack(peerConnection *webrtc.PeerConnection, currentUserID string) {

	_, err := peerConnection.AddTransceiver(webrtc.RTPCodecTypeVideo)
	errorHelper.CheckPanic(err)

	// Set a handler for when a new remote track starts, this just distributes all our packets
	// to connected peers
	peerConnection.OnTrack(func(remoteTrack *webrtc.Track, receiver *webrtc.RTPReceiver) {
		go callPeerConnection(remoteTrack, peerConnection)

		// Create a local track, all our SFU clients will be fed via this track
		// main track of the broadcaster
		localTrack, newTrackErr := peerConnection.NewTrack(remoteTrack.PayloadType(), remoteTrack.SSRC(), "video", "pion")
		errorHelper.CheckPanic(newTrackErr)

		// the channel that will have the local track that is used by the sender
		// the localTrack needs to be fed to the receiver
		localTrackChan := make(chan *webrtc.Track, 1)
		localTrackChan <- localTrack
		if existingChan, ok := peerConnectionMap[currentUserID]; ok {
			// feed the exiting track from user with this track
			existingChan <- localTrack
		} else {
			peerConnectionMap[currentUserID] = localTrackChan
		}

		rtpBuf := make([]byte, 1400)
		for { // for publisher only
			i, readErr := remoteTrack.Read(rtpBuf)
			errorHelper.CheckPanic(readErr)

			_, err := localTrack.Write(rtpBuf[:i])
			// ErrClosedPipe means we don't have any subscribers, this is ok if no peers have connected yet
			if err != nil && err != io.ErrClosedPipe {
				log.Fatal(err)
			}
		}
	})
}

// Send a PLI on an interval so that the publisher is pushing a keyframe every rtcpPLIInterval
// This can be less wasteful by processing incoming RTCP events, then we would emit a NACK/PLI when a viewer requests it
func callPeerConnection(remoteTrack *webrtc.Track, peerConnection *webrtc.PeerConnection) {
	ticker := time.NewTicker(rtcpPLIInterval)
	for range ticker.C {
		SSRC := remoteTrack.SSRC()
		pictureLoss := &rtcp.PictureLossIndication{MediaSSRC: SSRC}
		packets := []rtcp.Packet{pictureLoss}
		rtcpSendErr := peerConnection.WriteRTCP(packets)
		if rtcpSendErr != nil {
			fmt.Println(rtcpSendErr)
		}
	}
}

//Create a track using her peerConnection object that other user will get to listen UDP traffic
//The peerConnectMap is like a cache, the first peer save a track at channel to the second peer get this track


// user is the caller of the method
// if user connects before peer: create channel and keep listening till track is added
// if peer connects before user: channel would have been created by peer and track can be added by getting the channel from cache
func receiveTrack(peerConnection *webrtc.PeerConnection, peerID string) {
	_, ok := peerConnectionMap[peerID]
	if !ok {
		peerConnectionMap[peerID] = make(chan *webrtc.Track, 1)
	}
	localTrack := <-peerConnectionMap[peerID]
	peerConnection.AddTrack(localTrack)
}