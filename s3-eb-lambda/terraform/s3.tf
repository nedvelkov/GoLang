# S3 Bucket
resource "aws_s3_bucket" "bucket" {
  bucket = var.bucket_name
}

resource "aws_s3_bucket_notification" "on_change" {

  bucket = aws_s3_bucket.bucket.id
  lambda_function {

    lambda_function_arn = aws_lambda_function.go_lambda.arn
    events              = ["s3:ObjectCreated:*"]
    filter_prefix       = "for-process/"
    filter_suffix       = ".csv"
  }

}

resource "aws_lambda_permission" "s3_invoke_lambda_permission" {

  statement_id  = "AllowS3ToInvokeLambda"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.go_lambda.arn
  principal     = "s3.amazonaws.com"
  source_arn    = aws_s3_bucket.bucket.arn

}
