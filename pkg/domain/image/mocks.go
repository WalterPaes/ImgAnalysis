package image

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

var FakeUrl = "http://fakeurl.com"

type HttpConnectorSuccess struct{}

func (h HttpConnectorSuccess) DoGet(_ string) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(string(bytesResponse))),
	}, nil
}

type HttpConnectorFail struct{}

func (h HttpConnectorFail) DoGet(_ string) (*http.Response, error) {
	return nil, errors.New("an error")
}

type HttpConnectorFailWithWrongStatusCode struct{}

func (h HttpConnectorFailWithWrongStatusCode) DoGet(_ string) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       ioutil.NopCloser(bytes.NewBufferString("testing")),
	}, nil
}
