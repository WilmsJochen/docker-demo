package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main() {

	r := gin.Default()

	r.GET("/", Message)
	r.GET("/:any", Message)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func Message(c *gin.Context) {
    fmt.Println("Http request Received")
	c.String(http.StatusOK, "Go says: wE lOvE KUbErNetEs")
}
