# Prequirements

1.  Docker for desktop [Docker: Accelerated, Containerized Application Development](https://www.docker.com/)
2.  Localstack
    - Localstack CLI [Installation | Docs (localstack.cloud)](https://docs.localstack.cloud/getting-started/installation/#localstack-cli) (recommended)
    - Localstack with Docker-Compose [Installation | Docs (localstack.cloud)](https://docs.localstack.cloud/getting-started/installation/#docker-compose)
3.  AWS ClI [Installing or updating the latest version of the AWS CLI - AWS Command Line Interface (amazon.com)](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
4.  AWC CLI local [AWS Command Line Interface | Docs (localstack.cloud)](https://docs.localstack.cloud/user-guide/integrations/aws-cli/#:~:text=LocalStack%20AWS%20CLI%20%28awslocal%29%20awslocal%20is%20a%20thin,source%20code%20can%20be%20found%20on%20GitHub%3A%20https%3A%2F%2Fgithub.com%2Flocalstack%2Fawscli-local)
5.  Terraform [Install | Terraform | HashiCorp Developer](https://developer.hashicorp.com/terraform/downloads)
6.  Set Powershell _(in administrator mode)_ `Set-ExecutionPolicy -ExecutionPolicy RemoteSigned`
7.  Execute `go.exe install github.com/aws/aws-lambda-go/cmd/build-lambda-zip@latest`

# Test project

**Start docker and localstack.**

To start project execute script `deploy-terraform.ps1`.
For destroying resources execute script `destroy-terraform.ps1`

For testing dynamoDB you can use command awslocal dynamodb scan --table-name <table-name>. The command display all records in selected table.

For testing lambda function use command awslocal lambda invoke --function-name <function-name> <output-file> --payload file://<path-to-json-file> --cli-binary-format raw-in-base64-out, where output-file will save result from function, path-to-json-file for current project in folder lambda is provided invoke.json for testing api request.
