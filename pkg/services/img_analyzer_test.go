package services

import (
	"ImgAnalysis/internal/ports/analyzer"
	"ImgAnalysis/internal/ports/recognizer"
	"testing"
)

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
