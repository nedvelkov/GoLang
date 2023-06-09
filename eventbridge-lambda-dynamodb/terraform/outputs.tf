
output "id_gateway" {
  description = "Base URL for API Gateway stage."

  value = "http://localhost:4566/restapis/${aws_api_gateway_rest_api.rest.id}/test/_user_request_/test"
}
