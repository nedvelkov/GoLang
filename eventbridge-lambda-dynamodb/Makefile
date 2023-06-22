
# Variables
LAMBDA_NAME := go-lambda

.PHONY: build

help: ## This help.
	@cd ./lambda-dynamodb
	@echo "Test"


#commands
build:
	cd ./lambda-dynamodb && \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main

build-zip: build
	cd ./lambda-dynamodb && \
	build-lambda-zip -o ../build/main.zip main

update-lambda: build-zip
	awslocal lambda update-function-code --function-name $(LAMBDA_NAME) --zip-file fileb://build/main.zip

deploy-terraform: build-zip
	cd ./terraform && \
	terraform init && \
	terraform apply --auto-approve -var="lambda_name=$(LAMBDA_NAME)"
	

destroy-terraform:
	cd ./terraform && \
	terraform destroy --auto-approve -var="lambda_name=$(LAMBDA_NAME)"

