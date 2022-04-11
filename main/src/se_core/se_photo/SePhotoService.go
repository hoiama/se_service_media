package se_photo

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


//Get video by idUser and idVideo video folder
func getImageSvgFromServerByName(context *gin.Context) {
	image := context.Param("image")
	pathVideo := fmt.Sprintf("repository/image/svg/%s.svg", image)
	context.File(pathVideo)
	context.Header("Content-Type", "image/svg+xml")
}


//Get video by idUser and idVideo video folder
func getImagePngFromServerByName(context *gin.Context) {
	image := context.Param("image")
	pathVideo := fmt.Sprintf("repository/image/png/%s.png", image)
	context.File(pathVideo)
	context.Header("Content-Type", "image/png")
}