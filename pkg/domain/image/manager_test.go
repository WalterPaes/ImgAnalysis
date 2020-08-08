package image

import (
	"ImgAnalysis/internal/ports/net"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var fakeUrl = "http://fakeurl.com"

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

func TestNewImageManager(t *testing.T) {
	manager := NewImageManager(&HttpConnectorSuccess{})
	expected := "*image.Manager"
	got := fmt.Sprintf("%T", manager)
	if expected != got {
		t.Errorf("Was expected '%s' but got '%s'", expected, got)
	}
}

func TestManager_GetDataByUrl(t *testing.T) {
	getDataByUrl := func(connector net.HttpConnector) ([]byte, error) {
		manager := NewImageManager(connector)
		img := &ImageData{Url: fakeUrl}
		return manager.GetDataByUrl(img)
	}

	responseInBytes, err := getDataByUrl(&HttpConnectorSuccess{})
	if err != nil {
		t.Fatalf("Errors was not expected! Err: %s", err.Error())
	}

	t.Run("Assert Response Body", func(t *testing.T) {
		got := string(responseInBytes)
		expected := string(bytesResponse)

		if got != expected {
			t.Errorf("Was expected %v but got %v", expected, got)
		}
	})

	_, err = getDataByUrl(&HttpConnectorFail{})
	t.Run("Assert: Has Error", func(t *testing.T) {
		if err == nil {
			t.Errorf("An error was expected!")
		}
	})

	_, err = getDataByUrl(&HttpConnectorFailWithWrongStatusCode{})
	if err == nil {
		t.Fatal("An Error was expected!")
	}

	t.Run("Assert Error: Wrong Status Code", func(t *testing.T) {
		if err != ErrorInRequest {
			t.Errorf("An error was expected '%s'", ErrorInRequest.Error())
		}
	})
}
