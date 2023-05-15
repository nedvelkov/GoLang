package myFuncs

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type User struct {
	Email     string `json,dynamodbav:"email"`
	FirstName string `json,dynamodbav:"firstName"`
	LastName  string `json,dynamodbav:"lastName"`
}

var (
	ErrorMarshalling   = "Error marshalling user"
	ErrorUnmarshalling = "Error unmarshalling user"
	ErrorFetching      = "Error fetching user"
	ErrorCreating      = "Error creating user"
	ErrorAlreadyExists = "User already exists"
	ErrorInvalidEmail  = "Email is not valid"
	ErrorDeleting      = "Error deleting user"
)

func fetchRecord(email, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*User, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Email": {S: aws.String(email)}},
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

func createRecord(request events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*User, error) {
	var user User
	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		return nil, errors.New(ErrorUnmarshalling)
	}

	fetchUser, _ := fetchRecord(user.Email, tableName, dynamoClient)
	if len(fetchUser.Email) > 0 {
		return nil, errors.New(ErrorAlreadyExists)
	}

	// Marshall the user is not working correctly - add additional fields to the struct
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

func updateRecord(request events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*User, error) {
	var user User
	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		return nil, errors.New(ErrorUnmarshalling)
	}

	currentUser, _ := fetchRecord(user.Email, tableName, dynamoClient)
	if len(currentUser.Email) == 0 {
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

func deleteRecord(email, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Email": {S: aws.String(email)},
		},
		TableName: aws.String(tableName),
	}
	_, err := dynamoClient.DeleteItem(input)
	if err != nil {
		return errors.New(ErrorDeleting)
	}
	return nil
}
