package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/rpc/json"
	"net/http"
)

func main() {
	args := Args{"hellow"}
	msg, _ := json.EncodeClientRequest("HelloService.Hello", args)
	resp, _ := http.Post("http://localhost:9999/hi_rpc", "application/json", bytes.NewReader(msg))
	var result string
	json.DecodeClientResponse(resp.Body, &result)
	fmt.Println(result)
	}

type Args struct{
	H string
}
