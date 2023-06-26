& .\zip-main.ps1
awslocal lambda update-function-code --function-name go-lambda --zip-file fileb://../build/main.zip