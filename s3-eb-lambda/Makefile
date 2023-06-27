include .env
#environment variables
export BUCKET_NAME
export SQS_NAME

# Variables
LAMBDA_NAME := go-lambda

.PHONY: build

help: ## This help.
	@echo "Test"


#commands
build:
	powershell.exe ./scripts/build-main.ps1

build-zip: build
	powershell.exe ./scripts/zip-main.ps1

update-lambda: build-zip
	awslocal lambda update-function-code --function-name $(LAMBDA_NAME) --zip-file fileb://build/main.zip

upload-file:
	 awslocal s3 cp ./resources/list.csv  s3://$(BUCKET_NAME)/for-process/

get-logs:
	powershell.exe ./scripts/log-events.ps1 -lambdaName $(LAMBDA_NAME)

get-sqs-messages:
	awslocal sqs receive-message  --queue-url http://localhost:4566/000000000000/$(SQS_NAME) --attribute-names All --message-attribute-names All --max-number-of-messages 10

deploy-terraform:
	cd ./terraform && \
	terraform init && \
	terraform apply --auto-approve -var="lambda_name=$(LAMBDA_NAME)" \
	-var="bucket_name=$(BUCKET_NAME)" -var="sqs_name=$(SQS_NAME)"

destroy-terraform:
	awslocal s3 rm s3://$(BUCKET_NAME) --recursive
	awslocal s3 rm s3://$(EXPORT_BUCKET_NAME) --recursive
	cd ./terraform && \
	terraform destroy --auto-approve -var="lambda_name=$(LAMBDA_NAME)" \
	-var="bucket_name=$(BUCKET_NAME)" -var="sqs_name=$(SQS_NAME)"

