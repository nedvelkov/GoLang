module "eventbridge" {
  source = "terraform-aws-modules/eventbridge/aws"

  bus_name = var.eventbrdige_name

  rules = {
    lambda = {
      description   = "Capture all order data"
      event_pattern = jsonencode({ "Source" : ["test-event"] })
      enabled       = true
    }
  }

  targets = {
    lambda = [
      {
        name = "lambda-target"
        arn  = aws_lambda_function.go_lambda.arn
      },
    ]
  }

  tags = {
    Name = "my-bus"
  }
}

resource "aws_lambda_permission" "allow_eventbridge" {
  statement_id  = "AllowExecutionFromEventBridge"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.go_lambda.function_name
  principal     = "events.amazonaws.com"
}


# resource "aws_iam_role" "event_exec" {
#   name = "event_role_lambda"

#   assume_role_policy = jsonencode({
#     Version = "2012-10-17"
#     Statement = [{
#       Action = ["sts:AssumeRole"]
#       Effect = "Allow"
#       Sid    = ""
#       Principal = {
#         Service = "lambda.amazonaws.com"
#       }
#       },
#     ]
#   })
# }

# resource "aws_iam_role_policy_attachment" "event_policy" {
#   for_each = toset([
#     "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
#   ])

#   role       = aws_iam_role.event_exec.name
#   policy_arn = each.value
# }

