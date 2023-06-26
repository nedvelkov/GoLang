# module "eventbridge" {
#   source = "terraform-aws-modules/eventbridge/aws"

#   bus_name = var.eventbrdige_name

#   rules = {
#     lambda = {
#       description = "Capture creating object in s3 bucket"
#       event_pattern = jsonencode({
#         "Source" : ["aws.s3"],
#         "Detail-type" : ["Object Created"]
#         }
#       )
#       enabled = true
#     }
#   }

#   targets = {
#     lambda = [
#       {
#         name = "lambda-target"
#         arn  = aws_lambda_function.go_lambda.arn
#       },
#     ]
#   }

#   tags = {
#     Name = "my-bus"
#   }
# }

# resource "aws_lambda_permission" "allow_eventbridge" {
#   statement_id  = "AllowExecutionFromEventBridge"
#   action        = "lambda:InvokeFunction"
#   function_name = aws_lambda_function.go_lambda.function_name
#   principal     = "events.amazonaws.com"
# }



# resource "aws_cloudwatch_event_rule" "rule" {
#   name = "my-rule"
#   event_pattern = jsonencode({
#     source      = ["aws.s3"]
#     detail-type = ["AWS API Call via CloudTrail"]
#     detail = {
#       eventSource = ["s3.amazonaws.com"]
#       eventName   = ["PutObject"]
#       requestParameters = {
#         bucketName = [aws_s3_bucket.bucket.id]
#       }
#     }
#   })
# }

# resource "aws_cloudwatch_event_target" "target" {
#   rule      = aws_cloudwatch_event_rule.rule.name
#   target_id = aws_lambda_function.go_lambda.id
#   arn       = aws_lambda_function.go_lambda.arn
# }

# resource "aws_lambda_permission" "permission" {
#   statement_id  = "AllowExecutionFromCloudWatch"
#   action        = "lambda:InvokeFunction"
#   function_name = aws_lambda_function.go_lambda.arn
#   principal     = "events.amazonaws.com"
# }
