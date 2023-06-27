# Input variable definitions

variable "lambda_name" {
  description = "Name of lambda function"
  type        = string
  default     = "go-lambda"
}

variable "table_name" {
  description = "Name of dynamodb table"
  type        = string
  default     = "records"

}

variable "bucket_name" {
  default = "my-bucket"
  type    = string
}

variable "export_bucket" {
  default = "export-bucket"
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
