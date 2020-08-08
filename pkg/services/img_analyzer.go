package services

import (
	"ImgAnalysis/internal/ports/analyzer"
	"ImgAnalysis/internal/ports/recognizer"
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type ImgAnalyzer struct {
	svc recognizer.Recognizer
}

func NewAnalyzer(service recognizer.Recognizer) analyzer.Analyzer {
	return &ImgAnalyzer{svc: service}
}

func (a *ImgAnalyzer) DoAnalysis(img []byte) (*analyzer.Result, error) {
	// Create a specific Input to Rekognition
	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			// Set Img to Bytes
			Bytes: img,
		},
	}

	// Get Result of Rekognition
	output, err := a.svc.DetectLabels(input)
	if err != nil {
		return nil, err
	}

	// Marshal the output and get bytes
	outputInBytes, err := json.Marshal(output)
	if err != nil {
		return nil, err
	}

	// Set the bytes of output to specific struct
	var analysisResult *analyzer.Result
	err = json.Unmarshal(outputInBytes, &analysisResult)
	if err != nil {
		return nil, err
	}

	return analysisResult, nil
}
