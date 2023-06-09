package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (*events.APIGatewayProxyResponse, error) {
	return apiResponse(200, "Hello from Lambda")
}

func main() {
	lambda.Start(HandleRequest)
}

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	stringBody, _ := json.Marshal(body)
	response := events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: status,
		Body:       string(stringBody),
	}
	return &response, nil
}
