resource "aws_lambda_function" "hello_world" {
  function_name = "HelloWorld"

  filename = "${path.module}/lambda/main.zip"

  runtime = "go1.x"
  handler = "main"

  role = aws_iam_role.lambda_exec.arn
}

resource "aws_iam_role" "lambda_exec" {
  name = "serverless_lambda"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = ["sts:AssumeRole"]
      Effect = "Allow"
      Sid    = ""
      Principal = {
        Service = "lambda.amazonaws.com"
      }
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  for_each=toset([
    "arn:aws:iam::aws:policy/service-role/AWSLambdaDynamoDBExecutionRole",
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  ])

  role         = aws_iam_role.lambda_exec.name
  policy_arn   = each.value
  }
