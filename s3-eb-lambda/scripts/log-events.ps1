param(
    [Parameter(Mandatory = $true)]
    [string]$lambdaName,
    [int]$logsCount = 5
)


# Run AWS CLI command and capture the JSON response
$jsonResponse = & awslocal logs describe-log-streams --log-group-name /aws/lambda/$lambdaName --order-by LastEventTime --descending --limit $logsCount | Out-String

# Convert JSON to PowerShell object
$result = ConvertFrom-Json $jsonResponse

# Extract log streams from the JSON response
$logStreams = $result.logStreams

# Loop through each log stream and retrieve log events
foreach ($logStream in $logStreams) {
    $logStreamName = $logStream.logStreamName
    $logTime = (Get-Date 01.01.1970).AddSeconds($logStream.creationTime / 1000)  
    Write-Output "Log time: $($logTime)"
    Write-Output "Events:"

    # Run AWS CLI command to get log events for the log stream
    $jsonEventsResponse = & awslocal logs get-log-events --log-group-name /aws/lambda/$lambdaName --log-stream-name $logStreamName | Out-String

    # Convert JSON to PowerShell object
    $eventsResult = ConvertFrom-Json $jsonEventsResponse

    # Extract events from the JSON response
    $events = $eventsResult.events

    # Loop through each event and display the desired output format
    foreach ($event in $events) {
        $timestamp = $event.timestamp / 1000
        $date2 = (Get-Date 01.01.1970).AddSeconds($timestamp)  
        $message = $event.message
    
        Write-Host ("`t{0} : {1}" -f $date2, $message )
    }

    Write-Output ""
}