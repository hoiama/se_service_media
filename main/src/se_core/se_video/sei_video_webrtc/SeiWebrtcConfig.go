package sei_video_webrtc

import (
	//"sei_service_video/main/src/sei_entity/web_rtc"
	"time"
)

const (
	rtcpPLIInterval = time.Second * 3
)
//
//var peerConnectionMap = make(map[string]chan *webrtc.Track)
//var sdpWebRtcMap = make(map[string] web_rtc.SeWebrtcSdp)
//
//func configWebRTC()  {
//	configApiCodec()
//	configWebrtc()
//}
//
//
////Here we register the codec VP8
//func configApiCodec() *webrtc.API{
//	mediaEngine := webrtc.MediaEngine{}
//	codecVP8 := webrtc.NewRTPVP8Codec(webrtc.DefaultPayloadTypeVP8, 90000)
//	mediaEngine.RegisterCodec(codecVP8)
//	withMediaEngine := webrtc.WithMediaEngine(mediaEngine)
//	return webrtc.NewAPI(withMediaEngine)
//}
//
//func configWebrtc() webrtc.Configuration{
//	//Register the signaling server to discovery the ip and port to connect to
//	iceServer := webrtc.ICEServer{
//		URLs:           []string{"stun:stun.l.google.com:19302"},
//	}
//
//	iceServes := []webrtc.ICEServer{iceServer}
//
//	return webrtc.Configuration{
//		ICEServers:           iceServes,
//		ICETransportPolicy:   0,
//		BundlePolicy:         0,
//		RTCPMuxPolicy:        0,
//		PeerIdentity:         "",
//		Certificates:         nil,
//		ICECandidatePoolSize: 0,
//		SDPSemantics:         0,
//	}
//}
