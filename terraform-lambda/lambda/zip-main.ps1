#Script is based on https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html#golang-package-windows
#it must be executed every time lambda is updated!
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
go build -o main main.go
~\Go\Bin\build-lambda-zip.exe -o main.zip main
