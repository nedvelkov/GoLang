param(
    [string]$ApiId
)

$url = "http://localhost:4566/restapis/" + $ApiId + "/test/_user_request_/test"
$urlConstant = "const API_GATEWAY_URL = """ + $url + """;"

Set-Content -Path "./website/www/js/api.js" -Value $urlConstant

docker build ./website -t dockernginx
docker container run -d -p 5000:80 dockernginx:latest --name "website"
