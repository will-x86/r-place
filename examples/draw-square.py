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
