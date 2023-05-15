package main

import (
	"fmt"
	"lambda-test/myFuncs"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	dynamoClient dynamodbiface.DynamoDBAPI
	tableName    string
)

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
	tableName = "users"
	lambda.Start(HandleLambdaEvent)
}

func HandleLambdaEvent(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	switch request.HTTPMethod {
	case "GET":
		return myFuncs.GetUser(request, tableName, dynamoClient)
	case "POST":
		return myFuncs.CreateUser(request, tableName, dynamoClient)
	case "PUT":
		return myFuncs.UpdateUser(request, tableName, dynamoClient)
	case "DELETE":
		return myFuncs.DeleteUser(request, tableName, dynamoClient)
	default:
		return myFuncs.UnhandledMethod()
	}
}
