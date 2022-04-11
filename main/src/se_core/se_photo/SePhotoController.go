package se_photo

import (
	"github.com/gin-gonic/gin"
	"se_service_media/main/src/se_configuration/sei_router"
)

func SeMediaController(routerGin *gin.Engine)  {

	//Get a video by idUser and idVideo
	routerGin.GET(sei_router.ApiImage+ "svg/:image",
		func(context *gin.Context) {
			getImageSvgFromServerByName(context)
		},
	)

	//Get a video by idUser and idVideo
	routerGin.GET(sei_router.ApiImage+ "png/:image",
		func(context *gin.Context) {
			getImagePngFromServerByName(context)
		},
	)
}