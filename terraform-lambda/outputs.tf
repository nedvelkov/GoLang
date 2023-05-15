output "id_gateway" {
  description = "Base URL for API Gateway stage."

  value = aws_api_gateway_rest_api.rest.id
}

output "domain" {
  description = "Domain name of the bucket"
  value       = aws_s3_bucket_website_configuration.s3_bucket.website_domain
}

output "website_endpoint" {
  value = aws_s3_bucket_website_configuration.s3_bucket.website_endpoint
}