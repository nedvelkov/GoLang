awslocal dynamodb create-table --cli-input-json file://table-definition.json
awslocal dynamodb batch-write-item --request-items "file://test-items.json"
awslocal iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json
awslocal iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaDynamoDBExecutionRole
awslocal lambda create-function --function-name go-lambda-function --runtime go1.x --role arn:aws:iam::618836716276:role/lambda-ex --handler main --zip-file fileb://main.zip