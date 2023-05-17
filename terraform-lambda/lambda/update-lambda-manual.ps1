& .\zip-main.ps1
awslocal lambda update-function-code --function-name GoLambda --zip-file fileb://main.zip

#Invoke lambda with /from terraform-lambda directory
#awslocal lambda invoke --function-name GoLambda ..\response2.json --payload file://invoke.json --cli-binary-format raw-in-base64-out 
