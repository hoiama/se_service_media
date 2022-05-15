package se_video

import (
	"github.com/gin-gonic/gin"
)

func SeiSseController(routerGin *gin.Engine)  {

	////Register a SSE of user, used to notify user's frontend
	//routerGin.GET(sei_router.API + "webrtc/sse/user/:idUser",
	//	func(context *gin.Context) {
	//		idUser := context.Param("idUser")
	//		chanStream := make(chan web_rtc.SeWebrtcSdp)
	//		registerSseUser(idUser, chanStream)
	//		context.Stream(func(w io.Writer) bool {
	//			msg, isOpen := <-chanStream
	//			if isOpen {
	//				context.SSEvent("message", msg)
	//				return true
	//			}
	//			return false
	//		})
	//	},
	//)
	//
	//
	////Put a SDP in a SSE idUser, used to notify user's frontend
	//routerGin.POST(sei_router.API + "webrtc/sse/sender",
	//	func(context *gin.Context) {
	//		var sdpWebRtc web_rtc.SeWebrtcSdp
	//		err := context.ShouldBindJSON(&sdpWebRtc)
	//		chanStreamUser := getSseUser(sdpWebRtc.PeerTo)
	//		if err != nil {
	//			context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	//			return
	//		}
	//		chanStreamUser <- sdpWebRtc
	//	},
	//)
}

