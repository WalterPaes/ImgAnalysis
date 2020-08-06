package image

import (
	"io/ioutil"
	"net/http"
)

type Image struct {
	Url string `json:"url"`
}

func (img *Image) GetDataByUrl() ([]byte, error) {
	// Do Http request
	res, err := http.Get(img.Url)
	if err != nil {
		return nil, err
	}

	// Close connection in the end
	defer res.Body.Close()

	// Parse the img buffer to slice of bytes
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Return bytes of image
	return body, nil
}