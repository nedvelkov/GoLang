output "id_gateway" {
  description = "Base URL for API Gateway stage."

  value = aws_api_gateway_rest_api.rest.id
}
