import http.client

conn = http.client.HTTPSConnection("will-x86.com")

headers = {
    'User-Agent': "curl/7.68.0",
    'Accept': "*/*"
    }

conn.request("GET", "/api/pixelsq?x=1&y=2&hex=%23102932", headers=headers)

res = conn.getresponse()
data = res.read()

print(data.decode("utf-8"))
