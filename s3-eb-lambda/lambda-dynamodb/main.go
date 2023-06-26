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
	"net/url"
	"os"
	"path/filepath"
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

type Record struct {
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
	for _, record := range s3Event.Records {
		createRecord(tableName, record, dynamoClient)
	}

	return apiResponse(200, "Done")
}

func createRecord(tableName string, s3Event events.S3EventRecord, dynamoClient dynamodbiface.DynamoDBAPI) (*Record, error) {
	s3 := s3Event.S3

	if filepath.Ext(s3.Object.Key) != ".csv" {
		return nil, fmt.Errorf("file is not csv")
	}
	record := new(Record)
	record.Id = getGuid()
	record.Invoke = time.Now().Local().String()

	val, err := getS3Object(s3)

	if err != nil {
		record.Bucket = err.Error()
	} else {
		record.Bucket = val
		CopyObject(s3)
	}

	av, err := dynamodbattribute.MarshalMap(record)
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

	return record, nil
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

func getS3Object(e events.S3Entity) (string, error) {
	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(e.Bucket.Name),
		Key:    aws.String(e.Object.Key),
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

func getGuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func CopyObject(e events.S3Entity) error {
	source := e.Bucket.Name + "/" + e.Object.Key
	_, err := svc.CopyObject(&s3.CopyObjectInput{Bucket: aws.String("export-bucket/export/"),
		CopySource: aws.String(url.QueryEscape(source)), Key: aws.String(e.Object.Key)})
	if err != nil {
		return err
	}

	// Wait to see if the item got copied
	err = svc.WaitUntilObjectExists(&s3.HeadObjectInput{Bucket: aws.String("export-bucket/export/"), Key: aws.String(e.Object.Key)})
	if err != nil {
		return err
	}

	return err
}
