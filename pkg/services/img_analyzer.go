package services

import (
	"ImgAnalysis/pkg/domain/analyzer"
	"ImgAnalysis/pkg/domain/recognizer"
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type ImgAnalyzer struct {
	svc recognizer.Recognizer
}

type Result struct {
	Labels []struct {
		Confidence float64       `json:"Confidence"`
		Instances  []interface{} `json:"Instances"`
		Name       string        `json:"Name"`
	} `json:"Labels"`
}

func NewAnalyzer(service recognizer.Recognizer) analyzer.Analyzer {
	return &ImgAnalyzer{svc: service}
}

func (a *ImgAnalyzer) DoAnalysis(img []byte) (*Result, error) {
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
	bytes, err := json.Marshal(output)
	if err != nil {
		return nil, err
	}

	// Unmarshall the bytes to result to send specific data to client
	var result *Result
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
