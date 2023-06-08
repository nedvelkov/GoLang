awslocal iam create-role --role-name go-lambda-execution-role --assume-role-policy-document file://role-policy.json
awslocal iam attach-role-policy --role-name go-lambda-execution-role --policy-arn arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess

awslocal lambda create-function --function-name go-lambda --runtime go1.x --role arn:aws:iam::000000000000:role/go-lambda-execution-role --handler main --zip-file fileb://..\main.zip