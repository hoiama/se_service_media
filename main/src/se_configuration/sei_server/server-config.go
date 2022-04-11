package sei_server

import (
	"github.com/gin-gonic/gin"
	"se_service_media/main/src/se_core/se_photo"
)

//SeServerMedia Server Config a sei_server gin
func SeServerMedia() {
	routerGin := gin.Default()
	routerGin.Use(CORSMiddleware)
	se_photo.SeMediaController(routerGin)
	routerGin.Run(":8005")
}
