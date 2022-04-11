package errorHelper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//CheckError Throw a sei_error if some problem
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}


func FmtError(err error) {
	if err != nil {
		fmt.Println(err)
	}

}


func CheckErrorContext(err error, context *gin.Context, msg string) {
	if err != nil {
		context.JSON(http.StatusBadRequest, msg)
	}
}