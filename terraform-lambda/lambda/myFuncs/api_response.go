package myFuncs

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	stringBody, _ := json.Marshal(body)
	response := events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: status,
		Body:       string(stringBody),
	}
	return &response, nil
}
