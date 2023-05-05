package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func HandleLambdaEvent(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	// var user User

	// if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
	// 	return nil, errors.New("invalid json")
	// }
	message := fmt.Sprint("Hello from %v", request.Body)

	response := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"},
		Body: string(message), StatusCode: 200}
	return &response, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
