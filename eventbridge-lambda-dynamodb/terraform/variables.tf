# Input variable definitions

variable "lambda_name" {
  description = "Name of lambda function"
  type        = string
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

variable "event_source" {
  default     = "Name of event source"
  type        = string
  description = "aws.apigateway"

}
