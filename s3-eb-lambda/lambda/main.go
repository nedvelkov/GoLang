package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rs/zerolog"
)

var (
	logger zerolog.Logger
	sess   *session.Session
)

func main() {
	region := os.Getenv("AWS_REGION")
	url := os.Getenv("AWS_ENDPOINT_URL")
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Endpoint:         aws.String(url),
			Region:           aws.String(region),
			S3ForcePathStyle: aws.Bool(true),
		},
	}))

	logger = zerolog.New(os.Stdout)

	logger.Info().Msg("invoke lambda")
	lambda.Start(HandleLambdaEvent)
}

func HandleLambdaEvent(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		readRecord(record)
	}
}

func readRecord(s3Event events.S3EventRecord) {

	logger.Info().Msg("create S3 client")
	svc := s3.New(sess)

	logger.Info().Msg("process event record")
	s3 := s3Event.S3

	val, err := getS3Object(s3, svc)

	if err != nil {
		logger.Error().Msg(err.Error())
	} else {
		CopyObject(s3, svc)
		SendSqs(val)
	}
}

func getS3Object(e events.S3Entity, svc *s3.S3) (string, error) {
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

func CopyObject(e events.S3Entity, svc *s3.S3) error {
	logger.Info().Msg("copy file to export folder")

	source := e.Bucket.Name + "/" + e.Object.Key
	processFolder := e.Bucket.Name + "/processed"
	files := strings.Split(e.Object.Key, "/")
	fileName := files[len(files)-1]

	_, err := svc.CopyObject(&s3.CopyObjectInput{Bucket: aws.String(processFolder),
		CopySource: aws.String(url.QueryEscape(source)), Key: aws.String(fileName)})
	if err != nil {
		return err
	}

	err = svc.WaitUntilObjectExists(&s3.HeadObjectInput{Bucket: aws.String(processFolder), Key: aws.String(e.Object.Key)})
	if err != nil {
		return err
	}

	return err
}
