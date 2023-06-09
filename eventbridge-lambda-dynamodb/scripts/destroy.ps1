Set-Location .\terraform

terraform destroy --auto-approve

Set-Location ..

Remove-Item .\build -recurse
