package adapters

import (
	"ImgAnalysis/internal/ports/analyzer"
	"ImgAnalysis/pkg/domain/image"
	"fmt"
)

type ImageRequestInput struct {
	ImageUrl string `json:"image_url"`
}

type ImageRequestOutput struct {
	Analysis []string `json:"analysis"`
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

	return adapter.setOutput(result.Labels), err
}

func (adapter *ImageRecognizeAdapter) setOutput(labels []analyzer.Labels) *ImageRequestOutput {
	output := &ImageRequestOutput{}
	for _, v := range labels {
		output.Analysis = append(
			output.Analysis,
			fmt.Sprintf("This image has %.2f%% chance being a %s", v.Confidence, v.Name))
	}
	return output
}
