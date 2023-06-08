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
