resource "aws_api_gateway_rest_api" "rest" {
  name = "AWS Lambda"
}

resource "aws_api_gateway_resource" "resource" {
  rest_api_id = aws_api_gateway_rest_api.rest.id
  parent_id   = aws_api_gateway_rest_api.rest.root_resource_id
  path_part   = "test"
}

resource "aws_api_gateway_method" "method" {
  rest_api_id   = aws_api_gateway_rest_api.rest.id
  resource_id   = aws_api_gateway_resource.resource.id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "integration" {
  type                    = "AWS"
  rest_api_id             = aws_api_gateway_rest_api.rest.id
  resource_id             = aws_api_gateway_resource.resource.id
  http_method             = aws_api_gateway_method.method.http_method
  integration_http_method = "POST"

  uri = aws_lambda_function.hello_world.invoke_arn

  request_templates = {
    "application/json" = "#set($allParams = $input.params())\n{\n\"body-json\" : $input.json(\"$\"),\n\"params\" : {\n#foreach($type in $allParams.keySet())\n    #set($params = $allParams.get($type))\n\"$type\" : {\n    #foreach($paramName in $params.keySet())\n    \"$paramName\" : \"$util.escapeJavaScript($params.get($paramName))\"\n        #if($foreach.hasNext),#end\n    #end\n}\n    #if($foreach.hasNext),#end\n#end\n},\n\"stage-variables\" : {\n#foreach($key in $stageVariables.keySet())\n\"$key\" : \"$util.escapeJavaScript($stageVariables.get($key))\"\n    #if($foreach.hasNext),#end\n#end\n},\n\"context\" : {\n    \"api-id\" : \"$context.apiId\",\n    \"api-key\" : \"$context.identity.apiKey\",\n    \"http-method\" : \"$context.httpMethod\",\n    \"stage\" : \"$context.stage\",\n    \"source-ip\" : \"$context.identity.sourceIp\",\n    \"user-agent\" : \"$context.identity.userAgent\",\n    \"request-id\" : \"$context.requestId\",\n    \"resource-id\" : \"$context.resourceId\",\n    \"resource-path\" : \"$context.resourcePath\"\n    }\n}\n"
  }
}

resource "aws_api_gateway_method_response" "response_200" {
  rest_api_id = aws_api_gateway_rest_api.rest.id
  resource_id = aws_api_gateway_resource.resource.id
  http_method = aws_api_gateway_method.method.http_method
  status_code = "200"
}

resource "aws_api_gateway_integration_response" "integration_response_200" {
  rest_api_id = aws_api_gateway_rest_api.rest.id
  resource_id = aws_api_gateway_resource.resource.id
  http_method = aws_api_gateway_method.method.http_method
  status_code = aws_api_gateway_method_response.response_200.status_code

  response_parameters = {
    "method.response.header.Access-Control-Allow-Origin" : "'*'",
    "method.response.header.Content-Type" : "'text/html'"
  }

   depends_on = [
    aws_api_gateway_integration.integration
  ]

  response_templates = {
    "text/html" : "$input.path('$')"
  }
}