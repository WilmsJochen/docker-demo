package router

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func OcrRouter(route *gin.Engine) {
	route.POST("/ocr", controller.Ocr)
}
