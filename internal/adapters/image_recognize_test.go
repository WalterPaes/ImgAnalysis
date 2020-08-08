package adapters

import (
	"ImgAnalysis/internal/ports/net"
	"ImgAnalysis/internal/ports/recognizer"
	"ImgAnalysis/pkg/domain/image"
	"ImgAnalysis/pkg/services"
	"fmt"
	"testing"
)

func TestNewImageRecognizeAdapter(t *testing.T) {
	manager := image.NewImageManager(&image.HttpConnectorSuccess{})
	service := services.NewAnalyzer(&services.RecognizerMockSuccess{})
	adapter := NewImageRecognizeAdapter(manager, service)
	expected := "*adapters.ImageRecognizeAdapter"
	got := fmt.Sprintf("%T", adapter)
	if expected != got {
		t.Errorf("Was expected '%s' but got '%s'", expected, got)
	}
}

func TestImageRecognizeAdapter_Recognize(t *testing.T) {
	getAdapter := func(connector net.HttpConnector, recognizer recognizer.Recognizer) *ImageRecognizeAdapter {
		manager := image.NewImageManager(connector)
		service := services.NewAnalyzer(recognizer)
		return NewImageRecognizeAdapter(manager, service)
	}

	input := &ImageRequestInput{ImageUrl: image.FakeUrl}

	adapter := getAdapter(&image.HttpConnectorSuccess{}, &services.RecognizerMockSuccess{})
	output, err := adapter.Recognize(input)
	if err != nil {
		t.Fatalf("Errors was not expected! Error: %s", err.Error())
	}

	t.Run("Assert Success: Correct Label", func(t *testing.T) {
		got := output.Analysis[0]
		expected := "This image has 97.42% chance being a Mammal"
		if got != expected {
			t.Errorf("Was expected '%s' but got '%s'", expected, got)
		}
	})

	adapter = getAdapter(&image.HttpConnectorSuccess{}, &services.RecognizerMockFail{})
	_, err = adapter.Recognize(input)

	t.Run("Assert With Recognizer Fail: Has Error", func(t *testing.T) {
		if err == nil {
			t.Error("An error was expected!")
		}
	})

	adapter = getAdapter(&image.HttpConnectorFail{}, &services.RecognizerMockSuccess{})
	_, err = adapter.Recognize(input)
	t.Run("Assert With Connector Fail: Has Error", func(t *testing.T) {
		if err == nil {
			t.Error("An error was expected!")
		}
	})
}
