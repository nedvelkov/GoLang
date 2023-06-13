param(
    [string]$ApiId
)

#Set up URL for api gateway
$url = "http://localhost:4566/restapis/" + $ApiId + "/test/_user_request_/test"
$urlConstant = "const API_GATEWAY_URL = """ + $url + """;"

Set-Content -Path "./website/www/js/api.js" -Value $urlConstant

#Upload website to docker container
docker build ./website -t dockernginx
docker run -d -p 5000:80 --name "website" -it dockernginx:latest
