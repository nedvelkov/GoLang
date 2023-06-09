$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
go build -o main main.go

if (-not (Test-Path ..\build)) {
    try {
        New-Item -Path ..\build -ItemType Directory -ErrorAction Stop | Out-Null #-Force
    }
    catch {
    }

}

~\Go\Bin\build-lambda-zip.exe -o ..\build\response.zip main
