import requests

payload = {"x": "0", "y": "0", "hex": "#FF69B4"}
r = requests.get("https://will-x86.com/api/pixelsq", params=payload)
print(r.text)
