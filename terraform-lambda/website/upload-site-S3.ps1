param(
    [string]$ApiId,
    [string]$S3Name
)

#Set up URL for api gateway
$url = "http://localhost:4566/restapis/$($ApiId)/test/_user_request_/test"
$urlConstant = "const API_GATEWAY_URL = ""$($url)"";"

Set-Content -Path "./website/www/js/api.js" -Value $urlConstant

#Upload website to s3 bucket
awslocal s3 sync "./website/www" s3://$S3Name