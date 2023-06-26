Set-Location ./lambda

if (-not (Test-Path ..\build)) {
    try {
        New-Item -Path ..\build -ItemType Directory -ErrorAction Stop | Out-Null #-Force
    }
    catch {
    }

}

~\Go\Bin\build-lambda-zip.exe -o ..\build\main.zip main

Set-Location ..