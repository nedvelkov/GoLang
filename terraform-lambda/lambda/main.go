package main

import (
	"errors"
	"fmt"

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

	//method := request.HTTPMethod
	message := "Hello from Go!"
	// if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
	// 	return nil, errors.New(request.RequestContext.HTTPMethod)
	// }
	// if len(user.FirstName) > 0 {
	// 	message = fmt.Sprintf("Hello %v %v", user.FirstName, user.LastName)
	// } else {
	// 	message = "Hello from Go!"
	// }

	response := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"},
		Body: message, StatusCode: 200}
	return &response, nil
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String("http://172.17.0.2:4566"),
		Region:   aws.String("us-east-1"),
	})
	if err != nil {
		return
	}

	dynamoClient = dynamodb.New(sess)
	lambda.Start(HandleLambdaEvent)
}
