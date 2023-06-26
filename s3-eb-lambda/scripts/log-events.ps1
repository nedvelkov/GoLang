# Run AWS CLI command and capture the JSON response
$jsonResponse = & awslocal logs describe-log-streams --log-group-name /aws/lambda/go-lambda --order-by LastEventTime --descending --limit 1 | Out-String

# Convert JSON to PowerShell object
$result = ConvertFrom-Json $jsonResponse

# Extract values
$logStream = $result.logStreams[0].logStreamName


# Write-Host $logStream


$jsonResponse = & awslocal logs get-log-events --log-group-name /aws/lambda/go-lambda --log-stream-name $logStream | Out-String
# Convert JSON to PowerShell object
$result = ConvertFrom-Json $jsonResponse

# Extract events from the JSON response
$events = $result.events

# Loop through each event and display the desired output format
foreach ($event in $events) {
    $message = $event.message

    Write-Host $message
}