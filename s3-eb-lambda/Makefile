# Variables
LAMBDA_NAME := go-lambda
BUCKET_NAME := my-bucket
EXPORT_BUCKET_NAME := export-bucket
TABLE_NAME := records

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
	 awslocal s3 cp ./resources/list.csv  s3://$(BUCKET_NAME)

scan-dynamodb:
	awslocal dynamodb scan --table-name $(TABLE_NAME)

get-logs:
	powershell.exe ./scripts/log-events.ps1 -lambdaName $(LAMBDA_NAME)

deploy-terraform:
	cd ./terraform && \
	terraform init && \
	terraform apply --auto-approve -var="lambda_name=$(LAMBDA_NAME)" \
	-var="bucket_name=$(BUCKET_NAME)" -var="table_name=$(TABLE_NAME)" \
	-var="export_bucket=$(EXPORT_BUCKET_NAME)"

destroy-terraform:
	awslocal s3 rm s3://$(BUCKET_NAME) --recursive
	awslocal s3 rm s3://$(EXPORT_BUCKET_NAME) --recursive
	cd ./terraform && \
	terraform destroy --auto-approve -var="lambda_name=$(LAMBDA_NAME)" \
	-var="bucket_name=$(BUCKET_NAME)" -var="table_name=$(TABLE_NAME)" \
	-var="export_bucket=$(EXPORT_BUCKET_NAME)"

