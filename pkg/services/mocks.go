package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/service/rekognition"
)

const (
	labelName       = "Mammal"
	labelConfidence = 97.41643524169922
)

var rekognitionPayloadSampleResponse = fmt.Sprintf("{\n\t\"Labels\": [{\n\t\t\"Confidence\": %v,\n\t\t\"Instances\": [],\n\t\t\"Name\": \"%s\",\n\t\t\"Parents\": [{\n\t\t\t\"Name\": \"Animal\"\n\t\t}]\n\t}]\n}", labelConfidence, labelName)

// Success Mock
type RecognizerMockSuccess struct{}

func (r RecognizerMockSuccess) DetectLabels(input *rekognition.DetectLabelsInput) (*rekognition.DetectLabelsOutput, error) {
	var output *rekognition.DetectLabelsOutput
	json.Unmarshal([]byte(rekognitionPayloadSampleResponse), &output)
	return output, nil
}

// Fail Mock
type RecognizerMockFail struct{}

func (r RecognizerMockFail) DetectLabels(_ *rekognition.DetectLabelsInput) (*rekognition.DetectLabelsOutput, error) {
	return nil, errors.New("an error")
}
