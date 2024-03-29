package user

import (
	"encoding/json"
	"errors"
	"go-serverless/my-packages/validators"

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

func FetchUser(email, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*User, error) {
	input := &dynamodb.GetItemInput{
		Key:       map[string]*dynamodb.AttributeValue{"email": {S: aws.String(email)}},
		TableName: aws.String(tableName),
	}
	result, err := dynamoClient.GetItem(input)
	if err != nil {
		return nil, errors.New("Error fetching user")
	}

	item := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)

	if err != nil {
		return nil, errors.New("Error unmarshalling user")
	}
	return item, nil
}

func FetchUsers(tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*[]User, error) {
	input := &dynamodb.ScanInput{TableName: aws.String(tableName)}
	result, err := dynamoClient.Scan(input)
	if err != nil {
		return nil, errors.New("Error fetching users")
	}

	items := new([]User)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, items)

	if err != nil {
		return nil, errors.New("Error unmarshalling users")
	}
	return items, nil
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*User, error) {
	var user User
	if err := json.Unmarshal([]byte(req.Body), &user); err != nil {
		return nil, errors.New("Error unmarshalling user")
	}
	if !validators.IsEmailValid(user.Email) {
		return nil, errors.New("Email is not valid")
	}

	currentUser, _ := FetchUser(user.Email, tableName, dynamoClient)
	if currentUser != nil && len(currentUser.Email) != 0 {
		return nil, errors.New("User already exists")
	}

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return nil, errors.New("Error marshalling user")
	}

	input := &dynamodb.PutItemInput{Item: av, TableName: aws.String(tableName)}
	_, err = dynamoClient.PutItem(input)
	if err != nil {
		return nil, errors.New("Error inserting user")
	}

	return &user, nil
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*User, error) {
	var user User
	if err := json.Unmarshal([]byte(req.Body), &user); err != nil {
		return nil, errors.New("Error unmarshalling user")
	}

	currentUser, _ := FetchUser(user.Email, tableName, dynamoClient)
	if currentUser != nil && len(currentUser.Email) == 0 {
		return nil, errors.New("User does not exists")
	}

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return nil, errors.New("Error marshalling user")
	}

	input := &dynamodb.PutItemInput{Item: av, TableName: aws.String(tableName)}
	_, err = dynamoClient.PutItem(input)
	if err != nil {
		return nil, errors.New("Error inserting user")
	}

	return &user, nil
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) error {
	email := req.QueryStringParameters["email"]
	input := &dynamodb.DeleteItemInput{Key: map[string]*dynamodb.AttributeValue{"email": {S: aws.String(email)}}, TableName: aws.String(tableName)}
	_, err := dynamoClient.DeleteItem(input)
	if err != nil {
		return errors.New("Error deleting user")
	}

	return nil
}
