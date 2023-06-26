package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3"
)

type User struct {
	Id     string `dynamodbav:"Id"`
	Bucket string `dynamodbav:"Bucket"`
	Invoke string `dynamodbav:"Invoke"`
}

var (
	dynamoClient dynamodbiface.DynamoDBAPI
	tableName    string
	svc          *s3.S3
)

func main() {
	region := os.Getenv("AWS_REGION")
	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(os.Getenv("AWS_ENDPOINT_URL")),
		Region:           aws.String(region),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return
	}

	svc = s3.New(sess)
	dynamoClient = dynamodb.New(sess)
	tableName = "records"
	lambda.Start(HandleLambdaEvent)
}

func HandleLambdaEvent(ctx context.Context, s3Event events.S3Event) (*events.APIGatewayProxyResponse, error) {
	record, err := createRecord(tableName, s3Event, dynamoClient)
	if err != nil {
		return apiResponse(500, err.Error())
	}
	// err = Save(record, "record")
	// if err != nil {
	// 	return apiResponse(500, err.Error())
	// }
	return apiResponse(200, record)
}

func createRecord(tableName string, s3Event events.S3Event, dynamoClient dynamodbiface.DynamoDBAPI) (*User, error) {
	user := new(User)
	user.Id = getGuid()
	user.Invoke = time.Now().Local().String()

	for _, record := range s3Event.Records {
		s3 := record.S3
		val, err := getS3Object(s3.Object.Key)
		if err != nil {
			user.Bucket = err.Error()
		} else {
			user.Bucket = fmt.Sprintf("[%s - %s] Bucket = %s, Key = %s, Val=%v \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key, val)
			//Save(user, "record")
		}
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

func getS3Object(objectKey string) (string, error) {
	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("my-bucket"),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return "", err
	}
	r := csv.NewReader(bufio.NewReader(resp.Body))
	stringSlice := []string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		stringSlice = append(stringSlice, record...)
	}

	return strings.Join(stringSlice, " ,"), nil
}

func Save(value interface{}, key string) error {
	// p, err := json.Marshal(value)
	// if err != nil {
	// 	return err
	// }
	input := &s3.PutObjectInput{
		Body:   strings.NewReader("Hello World!"),
		Bucket: aws.String("export-bucket"),
		Key:    aws.String(key),
	}
	_, err := svc.PutObject(input)
	if err != nil {
		return err
	}
	return nil
}

func getGuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
