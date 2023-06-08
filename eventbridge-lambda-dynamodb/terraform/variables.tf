# Input variable definitions

variable "lambda_name" {
  description = "Name of lambda function"
  type        = string
  default     = "go-lambda"
}

variable "table_name" {
  description = "Name of dynamodb table"
  type        = string
  default     = "users"

}

variable "eventbrdige_name" {
  description = "Name of custom eventbrdge"
  type        = string
  default     = "lambda-bridge"
}
