require 'uri'
require 'net/http'
require 'openssl'

url = URI("https://will-x86.com/api/pixelsq?x=1&y=2&hex=%23102932")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true
http.verify_mode = OpenSSL::SSL::VERIFY_NONE

request = Net::HTTP::Get.new(url)
request["User-Agent"] = 'curl/7.68.0'
request["Accept"] = '*/*'

response = http.request(request)
puts response.read_body
