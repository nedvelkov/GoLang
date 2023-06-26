package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rs/zerolog"
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
	logger       zerolog.Logger
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

	logger = zerolog.New(os.Stdout)
	svc = s3.New(sess)
	dynamoClient = dynamodb.New(sess)
	tableName = "records"
	logger.Info().Msg("Invoke lambda")
	lambda.Start(HandleLambdaEvent)
}

func HandleLambdaEvent(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		createRecord(tableName, record, dynamoClient)
	}
}

func createRecord(tableName string, s3Event events.S3EventRecord, dynamoClient dynamodbiface.DynamoDBAPI) (*Record, error) {
	logger.Info().Msg("Process event record")
	s3 := s3Event.S3

	if filepath.Ext(s3.Object.Key) != ".csv" {
		return nil, fmt.Errorf("file is not csv")
	}
	record := new(Record)
	record.Id = getGuid()
	record.Invoke = s3Event.EventTime.String()

	val, err := getS3Object(s3)

	if err != nil {
		logger.Error().Msg(err.Error())
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
		logger.Error().Msg(err.Error())
		return nil, errors.New(err.Error())
	}

	return record, nil
}

func getS3Object(e events.S3Entity) (string, error) {
	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(e.Bucket.Name),
		Key:    aws.String(e.Object.Key),
	})
	if err != nil {
		logger.Error().Msg(err.Error())
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
			logger.Error().Msg(err.Error())
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
	logger.Info().Msg("Copy file to export bucket")
	source := e.Bucket.Name + "/" + e.Object.Key
	_, err := svc.CopyObject(&s3.CopyObjectInput{Bucket: aws.String("export-bucket/export/"),
		CopySource: aws.String(url.QueryEscape(source)), Key: aws.String(e.Object.Key)})
	if err != nil {
		return err
	}

	err = svc.WaitUntilObjectExists(&s3.HeadObjectInput{Bucket: aws.String("export-bucket/export/"), Key: aws.String(e.Object.Key)})
	if err != nil {
		return err
	}

	return err
}
