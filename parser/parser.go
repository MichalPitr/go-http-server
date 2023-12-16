package parser

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/MichalPitr/go-http-server/types"
)

// TODO: Improve parsing
func Parse(request string) (types.HttpRequest, error) {
	req := strings.Split(request, "\r\n")
	if len(req) < 1 {
		return types.HttpRequest{}, fmt.Errorf("Malformed request")
	}
	reqHeader := strings.Split(req[0], " ")

	return types.HttpRequest{
		Method:   reqHeader[0],
		Path:     filepath.Clean(reqHeader[1]),
		Protocol: reqHeader[2],
	}, nil
}
