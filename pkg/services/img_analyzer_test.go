package services

import (
	"ImgAnalysis/pkg/domain/image"
	"ImgAnalysis/pkg/rekognition"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestImgAnalyzer_DoAnalysis(t *testing.T) {
	svc := NewAnalyzer(rekognition.NewRekognition())

	img := &image.ImageData{
		Url: "https://images.pexels.com/photos/1741205/pexels-photo-1741205.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=650&w=940",
	}

	resp, err := http.Get(img.Url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	svc.DoAnalysis(body)
}
