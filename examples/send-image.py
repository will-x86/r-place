import requests
from PIL import Image

image = Image.open("image.png")
print(f"Original size : {image.size}")

image_resized = image.resize((10, 10))
image_resized.save("image-resized.png")


files = {"image": open("image-resized.png", "rb")}
data = {"x": "0", "y": "0"}

response = requests.post("https://will-x86.com/api/image10x10", files=files, data=data)

print(response.text)
print(response.headers)
