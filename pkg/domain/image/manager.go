package image

import (
	"ImgAnalysis/internal/ports/net"
	"errors"
	"io/ioutil"
	"net/http"
)

var ErrorInRequest = errors.New("just status OK is expected")

type Manager struct {
	httpConn net.HttpConnector
}

func NewImageManager(conn net.HttpConnector) *Manager {
	return &Manager{httpConn: conn}
}

func (m *Manager) GetDataByUrl(img *ImageData) ([]byte, error) {
	// Do Http request
	res, err := m.httpConn.DoGet(img.Url)
	if err != nil {
		return nil, err
	}

	// Close connection in the end
	defer res.Body.Close()

	// Check status code from response
	if res.StatusCode != http.StatusOK {
		return nil, ErrorInRequest
	}

	// Parse the img buffer to slice of bytes
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Return bytes of image
	return body, nil
}
