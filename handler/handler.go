package handler

import (
	"fmt"
	"log"
	"os"

	"github.com/MichalPitr/go-http-server/types"
)

func HandleGet(req types.HttpRequest) []byte {
	if req.Path == "/" {
		req.Path = "/index.html"
	}
	f, err := os.ReadFile("./www" + req.Path)
	if err != nil {
		log.Println(err)
		return []byte("HTTP/1.1 404 NOT FOUND\r\n")
	}
	return []byte(fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\n%s\r\n", f))
}
