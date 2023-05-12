package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var dynamoClient dynamodbiface.DynamoDBAPI

func HandleLambdaEvent(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	email := request.QueryStringParameters["email"]
	var user User
	if len(email) > 0 {

		input := &dynamodb.GetItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"Email": {S: aws.String(email)}},
			TableName: aws.String("users"),
		}
		result, err := dynamoClient.GetItem(input)
		if err != nil {
			return nil, errors.New("in dynamoDB client")
		}
		err = dynamodbattribute.UnmarshalMap(result.Item, &user)
		if err != nil {
			return nil, errors.New("Error unmarshalling from dynamoDB")
		}
		responseDb := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"},
			Body: fmt.Sprintf("Greeting from dynamodb!Hello from %v %v", user.FirstName, user.LastName), StatusCode: 200}
		return &responseDb, nil
	}

	message := "Hello from Go!"
	response := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"},
		Body: message, StatusCode: 200}
	return &response, nil
}

func main() {
	url := os.Getenv("LOCALSTACK_HOSTNAME")
	region := os.Getenv("AWS_REGION")
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(fmt.Sprintf("http://%v:4566", url)),
		Region:   aws.String(region),
	})
	if err != nil {
		return
	}

	dynamoClient = dynamodb.New(sess)
	lambda.Start(HandleLambdaEvent)
}
