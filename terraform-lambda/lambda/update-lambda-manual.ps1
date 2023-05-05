go build -o main main.go
~\Go\Bin\build-lambda-zip.exe -o main.zip main
awslocal lambda update-function-code --function-name HelloWorld --zip-file fileb://main.zip

#Invoke lambda with /from terraform-lambda directory
#awslocal lambda invoke --function-name=HelloWorld response.json --payload file://lambda/invoke.json --cli-binary-format raw-in-base64-out 