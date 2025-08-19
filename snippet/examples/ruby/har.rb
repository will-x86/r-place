require 'uri'
require 'net/http'

url = URI("http://localhost:8081/api/pixels")

http = Net::HTTP.new(url.host, url.port)

request = Net::HTTP::Post.new(url)
request["Content-Type"] = 'application/json'
request["User-Agent"] = 'curl/7.68.0'
request["Accept"] = '*/*'
request.body = "{\"x\":1, \"y\":2, \"hex\":\"#102932\"}"

response = http.request(request)
puts response.read_body