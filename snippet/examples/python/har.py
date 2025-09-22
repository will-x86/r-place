import http.client

conn = http.client.HTTPSConnection("place.willx86.com")

payload = "{\"x\":1, \"y\":2, \"hex\":\"#102932\"}"

headers = {
    'Content-Type': "application/json",
    'User-Agent': "curl/7.68.0",
    'Accept': "*/*"
    }

conn.request("POST", "/api/pixels", payload, headers)

res = conn.getresponse()
data = res.read()

print(data.decode("utf-8"))
