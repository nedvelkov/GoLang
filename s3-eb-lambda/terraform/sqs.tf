resource "aws_sqs_queue" "terraform_queue_deadletter" {
  name = var.dlq_sqs_name
}

resource "aws_sqs_queue" "terraform_queue" {
  name = var.sqs_name
  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.terraform_queue_deadletter.arn
    maxReceiveCount     = 4
  })
}
