import http.client

conn = http.client.HTTPConnection("localhost:8081")

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