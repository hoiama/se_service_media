package sei_video_webrtc

//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/pion/webrtc/v2"
//	"net/http"
//	"sei_service_video/main/src/sei_entity/web_rtc"
//)

////Send the SPD description from host to guest recovery after
//	routerGin.POST(sei_router.API + "webrtc/offer",
//		func(context *gin.Context) {
//			var sdpWebRtc web_rtc.SeWebrtcSdp
//			err := context.ShouldBindJSON(&sdpWebRtc)
//			if err != nil {
//				context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
//				return
//			}
//			fmt.Print("offer SDP")
//			registerSdpOfferWebRtc(sdpWebRtc)
//			context.JSON(http.StatusOK, "")
//		},
//	)
//
//
//	//Get the SPD description to guest use and after answer
//	routerGin.GET(sei_router.API + "webrtc/:idMeeting",
//		func(context *gin.Context) {
//			idMeeting := context.Param("idMeeting")
//			sdpWebRtc := getRegisterSdpWebRtc(idMeeting)
//			context.JSON(http.StatusOK, sdpWebRtc)
//		},
//	)

//import (
//	"github.com/pion/webrtc/v2"
//	"sei_service_video/main/src/sei_entity/web_rtc"
//	errorHelper2 "sei_service_video/main/src/sei_helper/errorHelper"
//)
//
//
//func createAnswer(sdpWebRtc web_rtc.SeWebrtcSdp) webrtc.SessionDescription{
//	sessionDescriptionOffer := webrtc.SessionDescription{}
//	web_rtc.DecodeBase64(sdpWebRtc.Sdp, &sessionDescriptionOffer)
//
//	peerConnection, errorPeer := configApiCodec().NewPeerConnection(configWebrtc())
//	errorHelper2.CheckPanic(errorPeer)
//
//	if !sdpWebRtc.IsSender {
//		receiveTrack(peerConnection, sdpWebRtc.PeerTo)
//	} else {
//		createTrack(peerConnection, sdpWebRtc.PeerTo)
//	}
//
//	//Set session description of remote peer
//	peerConnection.SetRemoteDescription(sessionDescriptionOffer)
//
//	//Create answer
//	answer, err := peerConnection.CreateAnswer(nil)
//	errorHelper2.CheckPanic(err)
//
//	//Set local description and start UDP listeners
//	error := peerConnection.SetLocalDescription(answer)
//	errorHelper2.CheckPanic(error)
//	return answer
//}
//
////Register a SeWebrtcSdp with a idMeeting
//func registerSdpOfferWebRtc(sdpWebRtc web_rtc.SeWebrtcSdp){
//	sdpWebRtcMap[sdpWebRtc.IdMeeting] = sdpWebRtc
//}

////Get the SeWebrtcSdp registered with idMeeting
//func getRegisterSdpWebRtc(idMeeting string) web_rtc.SeWebrtcSdp {
//	return sdpWebRtcMap[idMeeting]
//}
