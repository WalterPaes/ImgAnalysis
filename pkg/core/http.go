package core

import (
	"ImgAnalysis/internal/ports/net"
	"net/http"
)

type Http struct{}

func NewHttpConnector() net.HttpConnector {
	return &Http{}
}

func (h *Http) DoGet(url string) (resp *http.Response, err error) {
	return http.Get(url)
}
