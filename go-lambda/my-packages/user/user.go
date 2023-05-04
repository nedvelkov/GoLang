package user

import (
	"encoding/json"
	"errors"
	"go-lambda/my-packages/validators"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var (
	ErrorMarshalling   = "Error marshalling user"
	ErrorUnmarshalling = "Error unmarshalling user"
	ErrorFetching      = "Error fetching user"
	ErrorCreating      = "Error creating user"
	ErrorInvalidEmail  = "Email is not valid"
)

func FetchUser(email, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*User, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {S: aws.String(email)}},
		TableName: aws.String(tableName),
	}
	result, err := dynamoClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFetching)
	}
	item := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorUnmarshalling)
	}
	return item, nil
}

func CreateUser(request events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*User, error) {
	var user User
	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		return nil, errors.New(ErrorUnmarshalling)
	}
	if !validators.IsEmailValid(user.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return nil, errors.New(ErrorMarshalling)
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = dynamoClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCreating)
	}

	return &user, nil
}
