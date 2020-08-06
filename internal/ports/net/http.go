package net

import "net/http"

type HttpConnector interface {
	DoGet(url string) (resp *http.Response, err error)
}
