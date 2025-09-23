$headers=@{}
$headers.Add("User-Agent", "curl/7.68.0")
$headers.Add("Accept", "*/*")
$response = Invoke-WebRequest -Uri 'https://will-x86.com/api/pixelsq?x=1&y=2&hex=%23102932' -Method GET -Headers $headers
