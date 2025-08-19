require "http/client"

url = "http://localhost:8081/api/pixels"
headers = HTTP::Headers{
  "Content-Type" => "application/json"
  "User-Agent" => "curl/7.68.0"
  "Accept" => "*/*"
}
reqBody = "{\"x\":1, \"y\":2, \"hex\":\"#102932\"}"

response = HTTP::Client.post url, headers: headers, body: reqBody
puts response.body