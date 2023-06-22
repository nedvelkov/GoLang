resource "aws_lambda_function" "go_lambda" {
  function_name = var.lambda_name

  filename = "../build/main.zip"

  runtime = "go1.x"
  handler = "main"

  role = aws_iam_role.lambda_exec.arn
}

resource "aws_lambda_function" "api_response_lambda" {
  function_name = "api-lambda-response"

  filename = "../build/response.zip"

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
  for_each = toset([
    "arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess",
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  ])

  role       = aws_iam_role.lambda_exec.name
  policy_arn = each.value
}

data "aws_iam_policy_document" "lambda_event_bridge" {
  count = 1
  statement {
    sid    = "AllowEventBridgePutEvents"
    effect = "Allow"
    resources = [
      "arn:aws:events:us-east-1:000000000000:event-bus/${var.eventbrdige_name}",
    ]

    actions = [
      "events:PutEvents",
    ]
  }
}

resource "aws_iam_policy" "lambda_event_bridge" {
  count  = 1
  policy = data.aws_iam_policy_document.lambda_event_bridge[count.index].json
}

resource "aws_iam_role_policy_attachment" "lambda_event_bridge" {
  count      = 1
  policy_arn = aws_iam_policy.lambda_event_bridge[count.index].arn
  role       = aws_iam_role.lambda_exec.name
}
