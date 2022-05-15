package sei_video_webrtc

import (
	"github.com/gin-gonic/gin"
)

func SeiWebRTCController(routerGin *gin.Engine)  {
	//
	////Send the SPD description from host to guest recovery after
	//routerGin.POST(sei_router.API + "webrtc/offer",
	//	func(context *gin.Context) {
	//		var sdpWebRtc web_rtc.SeWebrtcSdp
	//		err := context.ShouldBindJSON(&sdpWebRtc)
	//		if err != nil {
	//			context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	//			return
	//		}
	//		fmt.Print("offer SDP")
	//		registerSdpOfferWebRtc(sdpWebRtc)
	//		context.JSON(http.StatusOK, "")
	//	},
	//)
	//
	//
	////Get the SPD description to guest use and after answer
	//routerGin.GET(sei_router.API + "webrtc/:idMeeting",
	//	func(context *gin.Context) {
	//		idMeeting := context.Param("idMeeting")
	//		sdpWebRtc := getRegisterSdpWebRtc(idMeeting)
	//		context.JSON(http.StatusOK, sdpWebRtc)
	//	},
	//)
}

