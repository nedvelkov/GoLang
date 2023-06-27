# Input variable definitions

variable "lambda_name" {
  description = "Name of lambda function"
  type        = string
  default     = "go-lambda"
}

variable "bucket_name" {
  default = "my-bucket"
  type    = string
}

variable "sqs_name" {
  default = "sqs"
  type    = string
}


variable "dlq_sqs_name" {
  default = "dlq-sqs"
  type    = string
}
