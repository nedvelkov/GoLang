# Terraform configuration

resource "aws_s3_bucket" "s3_bucket" {
  bucket = var.bucket_name
  tags   = var.tags
}

resource "aws_s3_bucket_website_configuration" "s3_bucket" {
  bucket = aws_s3_bucket.s3_bucket.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "error.html"
  }

}

resource "aws_s3_bucket_acl" "s3_bucket" {
  bucket = aws_s3_bucket.s3_bucket.id
  acl    = "public-read"
}

# resource "aws_s3_object" "object_www" {
#   depends_on   = [aws_s3_bucket.s3_bucket]
#   for_each     = fileset("../www/", "*")
#   bucket       = var.bucket_name
#   key          = basename(each.value)
#   source = "../www/${each.value}"
#   etag = filemd5("../www/${each.value}")
#   content_type = "text/html"
#   acl          = "public-read"
# }

resource "null_resource" "remove_and_upload_to_s3" {
  depends_on   = [aws_s3_bucket.s3_bucket]

  provisioner "local-exec" {
    command = "awslocal s3 sync ../www s3://${aws_s3_bucket.s3_bucket.id}"
  }
}

# resource "aws_s3_object" "object_assets" {
#   depends_on = [aws_s3_bucket.s3_bucket]
#   for_each   = fileset(path.module, "assets/*")
#   bucket     = var.bucket_name
#   key        = each.value
#   source     = "${each.value}"
#   etag       = filemd5("${each.value}")
#   acl        = "public-read"
# }

resource "aws_s3_bucket_policy" "s3_bucket" {
  bucket = aws_s3_bucket.s3_bucket.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid       = "PublicReadGetObject"
        Effect    = "Allow"
        Principal = "*"
        Action    = "s3:GetObject"
        Resource = [
          aws_s3_bucket.s3_bucket.arn,
          "${aws_s3_bucket.s3_bucket.arn}/*",
        ]
      },
    ]
  })
}