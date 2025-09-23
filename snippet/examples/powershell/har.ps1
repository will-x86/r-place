$headers=@{}
$headers.Add("Content-Type", "application/json")
$headers.Add("User-Agent", "curl/7.68.0")
$headers.Add("Accept", "*/*")
$response = Invoke-WebRequest -Uri 'https://will-x86.com/api/pixels' -Method POST -Headers $headers -ContentType 'application/json' -Body '{"x":1, "y":2, "hex":"#102932"}'
