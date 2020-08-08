package services

import (
	"ImgAnalysis/internal/ports/analyzer"
	"ImgAnalysis/internal/ports/recognizer"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"testing"
)

const (
	labelName       = "Mammal"
	labelConfidence = 97.41643524169922
)

var (
	rekognitionPayloadSampleResponse = fmt.Sprintf("{\n\t\"Labels\": [{\n\t\t\"Confidence\": %v,\n\t\t\"Instances\": [],\n\t\t\"Name\": \"%s\",\n\t\t\"Parents\": [{\n\t\t\t\"Name\": \"Animal\"\n\t\t}]\n\t}]\n}", labelConfidence, labelName)
	//rekognitionPayloadSampleResponseWithError = fmt.Sprintf("{\n\t\"Labels\": [{\n\t\t\"Confidence\": %v,\n\t\t\"Instances\": [],\n\t\t\"Name\": \"%s\",\n\t\t\"Parents\": [{\n\t\t\t\"Name\": \"Animal\"\n\t\t}]\n\t}]\n}", labelName, labelConfidence)
)

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

func TestImgAnalyzer_DoAnalysis(t *testing.T) {
	getService := func(recognizer recognizer.Recognizer) (*analyzer.Result, error) {
		svc := NewAnalyzer(recognizer)
		return svc.DoAnalysis([]byte("TESTANDO"))
	}

	result, err := getService(&RecognizerMockSuccess{})
	if err != nil {
		t.Fatalf("Errors was not expected! Error: %s", err.Error())
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

	_, err = getService(&RecognizerMockFail{})
	t.Run("Assert with error", func(t *testing.T) {
		if err == nil {
			t.Error("An error was expected!")
		}
	})
}
