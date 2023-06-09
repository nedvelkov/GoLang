Set-Location .\lambda-dynamodb
& .\zip-main.ps1

Set-Location ..\apigateway-lambda
& .\zip-main.ps1

Set-Location ..\terraform

terraform init

terraform apply --auto-approve

Set-Location ..