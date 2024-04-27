package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func apiResponse(status int, body any) (*events.APIGatewayProxyResponse, error) {
	b, _ := json.Marshal(body)

	return &events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-type": "application/json",
		},
		StatusCode: status,
		Body:       string(b),
	}, nil
}
