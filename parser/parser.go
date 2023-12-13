package parser

import (
	"fmt"
	"path/filepath"
	"strings"
)

type HttpRequest struct {
	Method   string
	Path     string
	Protocol string
}

// TODO: Improve parsing
func Parse(request string) (HttpRequest, error) {
	req := strings.Split(request, "\r\n")
	if len(req) < 1 {
		return HttpRequest{}, fmt.Errorf("Malformed request")
	}
	reqHeader := strings.Split(req[0], " ")

	return HttpRequest{
		Method:   reqHeader[0],
		Path:     filepath.Clean(reqHeader[1]),
		Protocol: reqHeader[2],
	}, nil

}
