package rekognition

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

const region = "us-east-1"

func NewRekognition() *rekognition.Rekognition {
	sess := session.New(&aws.Config{
		Region: aws.String(region),
	})
	return rekognition.New(sess)
}
