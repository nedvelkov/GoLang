package myFuncs

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrorMethodNotFound = "Method not found"

type ErrorBody struct {
	ErrorMessage string `json:"error,omitempty"`
}

func GetUser(request events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	email := request.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := fetchUser(email, tableName, dynamoClient)
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{*aws.String(err.Error())})
		}
		return apiResponse(http.StatusOK, result)
	}
	return apiResponse(http.StatusBadRequest, ErrorBody{*aws.String("email is required")})
}

func CreateUser(request events.APIGatewayProxyRequest, tableName string, dynamoClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := createUser(request, tableName, dynamoClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{*aws.String(err.Error())})
	}

	return apiResponse(http.StatusCreated, result)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorBody{*aws.String(ErrorMethodNotFound)})
}
