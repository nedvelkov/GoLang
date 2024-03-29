awslocal lambda create-function `
  --function-name apigw-lambda `
  --runtime nodejs16.x `
  --handler lambda.apiHandler `
  --memory-size 128 `
  --zip-file fileb://function.zip `
  --role arn:aws:iam::111111111111:role/apigw

#Go lambda
awslocal lambda create-function `
  --function-name apigw-lambda `
  --runtime go1.x `
  --handler main `
  --zip-file fileb://main.zip `
  --role arn:aws:iam::111111111111:role/apigw

awslocal apigateway create-rest-api --name 'API Gateway Lambda integration'

awslocal apigateway get-resources --rest-api-id 8r5un0yne1

awslocal apigateway create-resource `
  --rest-api-id 8r5un0yne1 `
  --parent-id xezjb24fur `
  --path-part "{test}"

  awslocal apigateway put-method `
  --rest-api-id 8r5un0yne1 `
  --resource-id odp0sz7yse `
  --http-method GET `
  --request-parameters "method.request.path.test=true" `
  --authorization-type "NONE"

  awslocal apigateway put-integration `
  --rest-api-id 8r5un0yne1 `
  --resource-id odp0sz7yse `
  --http-method GET `
  --type AWS_PROXY `
  --integration-http-method POST `
  --uri arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:000000000000:function:apigw-lambda/invocations `
  --passthrough-behavior WHEN_NO_MATCH

  awslocal apigateway create-deployment `
  --rest-api-id 8r5un0yne1 `
  --stage-name test

  http://localhost:4566/restapis/8r5un0yne1/test/_user_request_/test
  http://localhost:4566/restapis/trbzxdy3ur/test/_user_request_/test



    awslocal apigateway put-integration `
  --rest-api-id 8r5un0yne1 `
  --resource-id odp0sz7yse `
  --http-method GET `
  --type AWS_PROXY `
  --integration-http-method POST `
  --uri arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:000000000000:function:go-lambda-function/invocations `
  --passthrough-behavior WHEN_NO_MATCH

#Create a resource for POST method
awslocal apigateway put-method `
  --rest-api-id 8r5un0yne1 `
  --resource-id odp0sz7yse `
  --http-method POST `
  --request-parameters "method.request.path.test=true" `
  --authorization-type "NONE"

awslocal apigateway put-integration `
  --rest-api-id 8r5un0yne1 `
  --resource-id odp0sz7yse `
  --http-method POST `
  --type AWS_PROXY `
  --integration-http-method POST `
  --uri arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:000000000000:function:go-2-lambda-function/invocations `
  --passthrough-behavior WHEN_NO_MATCH