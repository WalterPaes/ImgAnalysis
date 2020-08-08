package image

import (
	"ImgAnalysis/internal/ports/net"
	"fmt"
	"testing"
)

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
		img := &ImageData{Url: FakeUrl}
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
