package main

import (
	"go-lambda/my-packages/handlers"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var dynamoClient dynamodbiface.DynamoDBAPI

const tableName = "My-table"

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return
	}
	dynamoClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "GET":
		return handlers.GetUser(request, tableName, dynamoClient)
	case "POST":
		return handlers.CreateUser(request, tableName, dynamoClient)
	default:
		return handlers.UnhandledMethod()
	}
}
