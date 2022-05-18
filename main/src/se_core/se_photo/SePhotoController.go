package se_photo

import (
	"github.com/gin-gonic/gin"
	se_config "se_service_share/se_microservices"
)

func SeMediaController(routerGin *gin.Engine)  {

	//Get a video by idUser and idVideo
	routerGin.GET(se_config.SeServiceGatewayMedia.Path + "/image/svg/:image",
		func(context *gin.Context) {
			getImageSvgFromServerByName(context)
		},
	)


	//Get a video by idUser and idVideo
	routerGin.GET(se_config.SeServiceGatewayMedia.Path + "/image/png/:image",
		func(context *gin.Context) {
			getImagePngFromServerByName(context)
		},
	)
}