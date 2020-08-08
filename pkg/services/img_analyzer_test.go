package services

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"testing"
)

const (
	labelName       = "Mammal"
	labelConfidence = 97.41643524169922
)

var rekognitionPayloadSampleResponse = fmt.Sprintf("{\n\t\"Labels\": [{\n\t\t\"Confidence\": %v,\n\t\t\"Instances\": [],\n\t\t\"Name\": \"%s\",\n\t\t\"Parents\": [{\n\t\t\t\"Name\": \"Animal\"\n\t\t}]\n\t}]\n}", labelConfidence, labelName)

type RecognizerMock struct{}

func (r RecognizerMock) DetectLabels(_ *rekognition.DetectLabelsInput) (*rekognition.DetectLabelsOutput, error) {
	var output *rekognition.DetectLabelsOutput
	json.Unmarshal([]byte(rekognitionPayloadSampleResponse), &output)
	return output, nil
}

func TestImgAnalyzer_DoAnalysis(t *testing.T) {
	svc := NewAnalyzer(&RecognizerMock{})

	result, err := svc.DoAnalysis([]byte("TESTANDO"))
	if err != nil {
		t.Errorf("Errors was not expected! Error: %s", err.Error())
	}

	expectedConfidence := result.Labels[0].Confidence
	expectedName := result.Labels[0].Name

	t.Run("Assert Confidence", func(t *testing.T) {
		if expectedConfidence != labelConfidence {
			t.Errorf("Was expected '%f' but got '%f'", labelConfidence, expectedConfidence)
		}
	})

	t.Run("Assert Label's Name", func(t *testing.T) {
		if expectedName != labelName {
			t.Errorf("Was expected '%s' but got '%s'", labelName, expectedName)
		}
	})
}
