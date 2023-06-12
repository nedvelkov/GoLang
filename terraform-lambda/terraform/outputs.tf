output "id_gateway" {
  description = "Base URL for API Gateway stage."

  value = aws_api_gateway_rest_api.rest.id
}

output "s3_bucket_name" {
  description = "S3 bucket name for static deploy"

  value = aws_s3_bucket.s3_bucket.id
}

output "website_endpoint" {
  value = "http://${aws_s3_bucket.s3_bucket.id}.s3-website.localhost.localstack.cloud:4566"
}