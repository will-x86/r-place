import requests

r = requests.post(
    "https://will-x86.com/api/pixels", json={"x": 0, "y": 0, "hex": "#FFFFFF"}
)
print(r.text)
