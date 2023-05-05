awslocal dynamodb create-table --cli-input-json file://table-definition.json
awslocal dynamodb batch-write-item --request-items "file://test-items.json"
awslocal iam create-role --role-name "go-lambda-execution-role" --assume-role-policy-document file://execution-role.json
awslocal lambda create-function --function-name "go-lambda-function" --zip-file fileb://main.zip --handler main --runtime go1.x --role arn:aws:iam::000000000000:role/go-lambda-execution-role
