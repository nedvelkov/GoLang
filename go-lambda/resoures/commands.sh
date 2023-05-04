#Create dynamodb table execute command in same directory as table-definition.json
aws --endpoint-url=http://localhost:4566 dynamodb create-table --cli-input-json file://table-definition.json 


awslocal dynamodb batch-write-item --request-items "file://test-items.json"

#Create iam role execute command in same directory as execution-role.json
aws --endpoint-url=http://localhost:4566 iam create-role --role-name "go-lambda-execution-role" --assume-role-policy-document file://execution-role.json
awslocal iam create-policy-version --policy-arn arn:aws:iam::000000000000:role/go-lambda-execution-role --policy-document file://execution-role-v2.json --set-as-default

#Create zip file
#follow instruction in https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html

#Upload lambda function,  note the role arn and execute command in same directory as main.go
aws --endpoint-url=http://localhost:4566 lambda create-function --function-name "go-2-lambda-function" --zip-file fileb://main.zip --handler main --runtime go1.x --role arn:aws:iam::000000000000:role/go-lambda-execution-role #arn of the role created above

#Invoke lambda function execute command in same directory as apigateway-aws-proxy-post.json and select output.json save location
aws --endpoint-url=http://localhost:4566 lambda invoke --cli-binary-format raw-in-base64-out --function-name go-2-lambda-function --invocation-type RequestResponse --no-sign-request --payload file://apigateway-aws-proxy-post.json --endpoint-url=http://localhost:4566 ..\output.json
awslocal lambda invoke --cli-binary-format raw-in-base64-out --function-name go-lambda-function --invocation-type RequestResponse --no-sign-request ..\output.json

#Update lambda function,  note the role arn and execute command in same directory as main.go
aws --endpoint-url=http://localhost:4566 lambda update-function-code --function-name "go-2-lambda-function" --zip-file fileb://main.zip

#Create api gateway, using this tutorial https://docs.localstack.cloud/user-guide/aws/apigateway/
aws --endpoint-url=http://localhost:4566 apigateway create-rest-api --name 'API Gateway Lambda integration'
awslocal apigateway create-rest-api --name 'API Gateway Lambda integration'

#Get Id of api gateway - jid4k75zyv
aws apigateway --endpoint-url=http://localhost:4566 get-resources --rest-api-id jid4k75zyv
awslocal apigateway get-resources --rest-api-id r9ohfdf7tz

#Get Id of resource - 9nar5yvb6a
awslocal apigateway create-resource  --rest-api-id jid4k75zyv   --parent-id 9nar5yvb6a   --path-part "{users}"
awslocal apigateway create-resource  --rest-api-id r9ohfdf7tz   --parent-id kdzj1cr5tf   --path-part "{users}"

#Get Id of creating resource - 7fke7yrs77
awslocal apigateway put-method `
 --rest-api-id jid4k75zyv `
  --resource-id 7fke7yrs77 `
   --http-method GET  --request-parameters "method.request.path.users=true" --authorization-type "NONE"

awslocal apigateway put-method `
 --rest-api-id r9ohfdf7tz `
  --resource-id j5d64j1syz `
   --http-method GET  --request-parameters "method.request.path.users=true" --authorization-type "NONE"

#Get Id of creating resource - 7fke7yrs77
awslocal apigateway update-method `
 --rest-api-id jid4k75zyv `
  --resource-id 7fke7yrs77 `
   --http-method GET  --request-parameters "method.request.path.users=false" --authorization-type "NONE"

#Integrate  lambda function with api gateway
awslocal apigateway put-integration  --rest-api-id jid4k75zyv  --resource-id 7fke7yrs77  --http-method GET  --type AWS_PROXY   --integration-http-method POST `
  --uri arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:000000000000:function:go-2-lambda-function/invocations --passthrough-behavior WHEN_NO_MATCH

awslocal apigateway put-integration  --rest-api-id r9ohfdf7tz  --resource-id j5d64j1syz  --http-method GET  --type AWS_PROXY   --integration-http-method POST `
  --uri arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:000000000000:function:go-lambda-function/invocations `
  --passthrough-behavior WHEN_NO_MATCH

 #Integrate  lambda function with api gateway
awslocal apigateway update-integration  --rest-api-id jid4k75zyv  --resource-id 7fke7yrs77  --http-method GET `
 --patch-operations "uri='arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:000000000000:function:go-2-lambda-function/invocations'"

#Create deployment
awslocal apigateway create-deployment `
--rest-api-id jid4k75zyv `
--stage-name test

awslocal apigateway create-deployment `
--rest-api-id r9ohfdf7tz `
--stage-name test

#Test deployment
curl -X GET http://localhost:4566/restapis/jid4k75zyv/stages/dev/resources/users
curl -X GET http://localhost:4566/restapis/r9ohfdf7tz/stages/test/_user_request_/test


awslocal apigateway delete-integration  --rest-api-id jid4k75zyv  --resource-id 7fke7yrs77  --http-method GET
awslocal apigateway delete-integration  --rest-api-id r9ohfdf7tz  --resource-id j5d64j1syz  --http-method GET