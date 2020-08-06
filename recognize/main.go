package main

import (
	"ImgAnalysis/internal/adapters"
	"ImgAnalysis/pkg/core"
	"ImgAnalysis/pkg/domain/image"
	"ImgAnalysis/pkg/rekognition"
	"ImgAnalysis/pkg/services"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

func Handler(req Request) (Response, error) {
	var input *adapters.ImageRequestInput
	err := json.Unmarshal([]byte(req.Body), &input)
	if err != nil {
		return Response{StatusCode: http.StatusInternalServerError}, err
	}

	adapter := adapters.NewImageRecognizeAdapter(
		image.NewImageManager(core.NewHttpConnector()),
		services.NewAnalyzer(rekognition.NewRekognition()))

	output, err := adapter.Recognize(input)
	if err != nil {
		return Response{StatusCode: http.StatusInternalServerError}, err
	}

	body, err := json.Marshal(output)
	if err != nil {
		return Response{StatusCode: http.StatusInternalServerError}, err
	}

	resp := Response{
		StatusCode: http.StatusOK,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
