package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func SendSqs(m string) {
	logger.Info().Msg("create SQS client")
	svc := sqs.New(sess)

	queueURL, err := getQueueUrl("sqs", svc)
	if err != nil {
		logger.Error().Msg(err.Error())
		return
	}

	sendMessageInput := &sqs.SendMessageInput{
		MessageBody: aws.String(m),
		QueueUrl:    queueURL.QueueUrl,
	}

	logger.Info().Msg("sending message")
	result, err := svc.SendMessage(sendMessageInput)
	if err != nil {
		logger.Error().Msg(err.Error())
		return
	}

	logger.Info().Msg(fmt.Sprintf("send message with body %v", m))
	logger.Info().Msg(fmt.Sprintf("send message with id %v", *result.MessageId))
}

func getQueueUrl(queueName string, svc *sqs.SQS) (*sqs.GetQueueUrlOutput, error) {

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})

	return result, err
}
