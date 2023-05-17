param(
    [string]$ApiId,
    [string]$S3Name
)


$url = "http://localhost:4566/restapis/" + $ApiId + "/test/_user_request_/test"
$urlConstant = "const API_GATEWAY_URL = """ + $url + """;"
$sitePath=$MyInvocation.MyCommand.Path + "/www"

Set-Content -Path "./www/js/api.js" -Value $urlConstant

awslocal s3 sync "./www" s3://$S3Name
