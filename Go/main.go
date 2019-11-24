package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
    go sendHttpCall()

	r := gin.Default()

	r.GET("/", Message)

	r.Run() // listen and serve on 0.0.0.0:8080


}

func Message(c *gin.Context) {
	c.String(http.StatusOK, "wE lOvE KUbErNetEs")
}

func sendHttpCall() {
    url := os.Getenv("url")
    if url == "" {
        url = "http://localhost:8080"
    }
    fmt.Println("URL: ",url)

    resp, err := http.Get(url)
    if err != nil{
        fmt.Println("something went wrong: ", err)
    }
    defer resp.Body.Close()

    body, err :=  ioutil.ReadAll(resp.Body)
    if err != nil{
        fmt.Println("something went wrong reading body: ", err)
    }
    fmt.Println(string(body))
}
