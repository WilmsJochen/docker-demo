package main

import (
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
 sendHttpCall()
}


func sendHttpCall() {
    url := os.Getenv("URL")
    if url == "" {
        url = "http://localhost:8080"
    }
    fmt.Println("Request send to URL: ",url)

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
