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
