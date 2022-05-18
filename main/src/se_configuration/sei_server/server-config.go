package sei_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"se_service_media/main/src/se_core/se_photo"
	se_config "se_service_share/se_microservices"
)

//SeServerMedia Server Config a sei_server gin
func SeServerMedia() {
	engine := gin.Default()
	engine.Use(CORSMiddleware)
	se_photo.SeMediaController(engine)
	fmt.Println(">>> listening at " + se_config.SeServiceGatewayMedia.UrlHttp)
	engine.Run(se_config.SeServiceGatewayMedia.PortHttp)
}
