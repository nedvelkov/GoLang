# Variables
LAMBDA_NAME := go-lambda
BUCKET_NAME := my-bucket
TABLE_NAME := records

.PHONY: build

help: ## This help.
	@echo "Test"


#commands
build:
	cd ./lambda-dynamodb && \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 && \
	go build -o main main.go

build-zip: build
	cd ./lambda-dynamodb && \
	build-lambda-zip -o ../build/main.zip main

update-lambda:
	powershell.exe awslocal lambda update-function-code --function-name $(LAMBDA_NAME) --zip-file fileb://build/main.zip

invoke-lambda:
	powershell.exe awslocal lambda invoke --function-name $(LAMBDA_NAME) --invocation-type Event --payload file://resources/s3event.json response.json

scan-dynamodb:
	powershell.exe awslocal dynamodb scan --table-name $(TABLE_NAME)

deploy-terraform:
	cd ./terraform && \
	terraform init && \
	terraform apply --auto-approve -var="lambda_name=$(LAMBDA_NAME)" \
	-var="bucket_name=$(BUCKET_NAME)" -var="table_name=$(TABLE_NAME)"
	

destroy-terraform:
	powershell.exe awslocal s3 rm s3://$(BUCKET_NAME) --recursive
	cd ./terraform && \
	terraform destroy --auto-approve -var="lambda_name=$(LAMBDA_NAME)"

