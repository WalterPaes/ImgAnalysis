package recognizer

import "github.com/aws/aws-sdk-go/service/rekognition"

type Recognizer interface {
	DetectLabels(input *rekognition.DetectLabelsInput) (*rekognition.DetectLabelsOutput, error)
}
