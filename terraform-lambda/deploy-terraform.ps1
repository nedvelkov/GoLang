terraform init
terraform apply -auto-approve

$apiId= terraform output id_gateway
$bucket_name = terraform output s3_bucket_name

$apiId=$apiId.Replace('"','')

& .\website\upload-site-S3.ps1 $apiId $bucket_name

& .\website\upload-site-Docker.ps1 $apiId
