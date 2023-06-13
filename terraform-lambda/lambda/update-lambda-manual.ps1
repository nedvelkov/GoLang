#Build and zip lambda
& .\zip-main.ps1

#Update lambda in localstac. Script may be updated with parameter for lambda name (in the current situation, the name is hardcoded GoLambda)
awslocal lambda update-function-code --function-name GoLambda --zip-file fileb://main.zip

#Invoke lambda with /from terraform-lambda directory
#awslocal lambda invoke --function-name GoLambda ..\response2.json --payload file://invoke.json --cli-binary-format raw-in-base64-out 
