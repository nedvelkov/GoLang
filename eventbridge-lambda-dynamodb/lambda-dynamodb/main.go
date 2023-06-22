package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type User struct {
	Id        string `dynamodbav:"Id"`
	FirstName string `dynamodbav:"FirstName"`
	LastName  string `dynamodbav:"LastName"`
	Invoke    string `dynamodbav:"Invoke"`
}

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

func HandleLambdaEvent(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	record, err := createRecord(request.Body, tableName, dynamoClient)
	if err != nil {
		return apiResponse(500, err.Error())
	}
	return apiResponse(200, record)
}

func createRecord(body, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*User, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return nil, errors.New("Error Generating UUID")
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	user, err := unmarshalJsonString(body)

	if user == nil {
		user = &User{}
	}

	if len(user.FirstName) == 0 {
		user.FirstName = "John"
		user.LastName = "Doe"
	}

	user.Id = uuid
	if err != nil {
		user.Invoke = err.Error()
	} else {
		user.Invoke = time.Now().Local().GoString()
	}

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return nil, errors.New("ErrorMarshalling")
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = dynamoClient.PutItem(input)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return user, nil
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

func unmarshalJsonString(text string) (*User, error) {
	var user User

	if err := json.Unmarshal([]byte(text), &user); err != nil {
		fmt.Println(err)
		return nil, errors.New(text)
	}
	return &user, nil
}
