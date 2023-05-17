param(
    [string]$ApiId,
    [string]$S3Name
)

$url = "http://localhost:4566/restapis/$($ApiId)/test/_user_request_/test"
$urlConstant = "const API_GATEWAY_URL = ""$($url)"";"

Set-Content -Path "./website/www/js/api.js" -Value $urlConstant

awslocal s3 sync "./website/www" s3://$S3Name