output "function_name" {
  description = "Name of the Lambda function."

  value = aws_lambda_function.hello_world.function_name
}

output "id_gateway" {
  description = "Base URL for API Gateway stage."

  value = aws_api_gateway_rest_api.rest.id
}
