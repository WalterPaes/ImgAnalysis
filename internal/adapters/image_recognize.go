package adapters

import (
	"ImgAnalysis/internal/ports/analyzer"
	"ImgAnalysis/pkg/domain/image"
	"encoding/json"
)

type ImageRequestInput struct {
	ImageUrl string `json:"image_url"`
}

type ImageRequestOutput struct {
	Labels []struct {
		Confidence float64       `json:"Confidence"`
		Instances  []interface{} `json:"Instances"`
		Name       string        `json:"Name"`
	} `json:"Labels"`
}

type ImageRecognizeAdapter struct {
	manager  *image.Manager
	analyzer analyzer.Analyzer
}

func NewImageRecognizeAdapter(manager *image.Manager, analyzer analyzer.Analyzer) *ImageRecognizeAdapter {
	return &ImageRecognizeAdapter{
		manager:  manager,
		analyzer: analyzer,
	}
}

func (adapter *ImageRecognizeAdapter) Recognize(req *ImageRequestInput) (*ImageRequestOutput, error) {
	// Init the Image Model with Request Data
	img := &image.ImageData{Url: req.ImageUrl}

	// Get the image bytes
	imgBytes, err := adapter.manager.GetDataByUrl(img)
	if err != nil {
		return nil, err
	}

	// Do the image analysis
	result, err := adapter.analyzer.DoAnalysis(imgBytes)
	if err != nil {
		return nil, err
	}

	// Return the Result of analysis to Client
	var output *ImageRequestOutput
	err = json.Unmarshal(result, output)
	if err != nil {
		return nil, err
	}

	return output, err
}
