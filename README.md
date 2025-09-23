# Examples 

Set a single pixel :) 

```python
import requests

payload = {"x": "0", "y": "0", "hex": "#FF69B4"}
r = requests.get("https://will-x86.com/api/pixelsq", params=payload)
print(r.text)
```

Set a single pixel (json)

```python
import requests

r = requests.post(
    "https://will-x86.com/api/pixels", json={"x": 0, "y": 0, "hex": "#FFFFFF"}
)
print(r.text)
```



Draw a square ! 
```python
import requests


# Function to set a pixel
def set_pixel(x, y, color):
    r = requests.post(
        "https://will-x86.com/api/pixels", json={"x": x, "y": y, "hex": color}
    )
    print(r.text)


color = "#FF0000"
for x in range(5, 10):
    for y in range(5, 10):
        set_pixel(x, y, color)

```


Print an image :) 
```python 
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
```


A smiley face!
```python
import requests

def set_pixel(x, y, color):
    requests.post("https://will-x86.com/api/pixels", json={"x": x, "y": y, "hex": color})

smiley = [
    [0,0,1,1,1,1,0,0],
    [0,1,0,0,0,0,1,0],
    [1,0,1,0,1,0,1,0],
    [1,0,0,0,0,0,1,0],
    [1,0,1,0,0,1,0,0],
    [1,0,0,1,1,0,0,0],
    [0,1,0,0,0,0,1,0],
    [0,0,1,1,1,1,0,0]
]

for y, row in enumerate(smiley):
    for x, pixel in enumerate(row):
        color = "#000000" if pixel else "#FFFFFF"
        set_pixel(x+5, y+5, color)
```
This is the code for an r-place clone that can only be used via code:)

Elements:
cmd/ -- entrypoint for golang
pkg/ -- packages for golang
ui/ -- React userinterface for requests
curl/ -- Basic curl commands for interfacing
.air.toml -- hot reloading

# Cache:

- K:V are stored in appendonly valkey db
- On startup all key:pair's are fetched up to xmax and ymax ( see .env.example )
- Two caches are stored, both the fully marshalled

# Development:

## Enviroment variables

Copy .env.example to .env and fillout values

## Nix

- Use nix package manager and run direnv allow

```
(cd ui && npm install && npm run build) && docker compose up -d && air
``

## non-nix
- Things required:
    - Air
    - Golang
    - docker compose
    - npm
```

(cd ui && npm install && npm run build) && docker compose up -d && air
``

## Windows

uhhhhhhhhhhh todo later

# Common errors

- Panics on startup
  - most likely missing an env variable
- Valkey keeps restarting constantly ( most likely
  - Most likely missed enviroment variables, ` docker compose down -v` then restart
