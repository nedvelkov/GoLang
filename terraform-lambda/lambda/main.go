package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func HandleLambdaEvent(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	var user User

	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		return nil, errors.New("invalid json")
	}

	response := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"},
		Body: fmt.Sprintf("Update via powershell!Hello from %v %v", user.FirstName, user.LastName), StatusCode: 200}
	return &response, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
