import requests


def set_pixel(x, y, color):
    requests.post(
        "https://will-x86.com/api/pixels", json={"x": x, "y": y, "hex": color}
    )


smiley = [
    [0, 0, 1, 1, 1, 1, 0, 0],
    [0, 1, 0, 0, 0, 0, 1, 0],
    [1, 0, 1, 0, 1, 0, 1, 0],
    [1, 0, 0, 0, 0, 0, 1, 0],
    [1, 0, 1, 0, 0, 1, 0, 0],
    [1, 0, 0, 1, 1, 0, 0, 0],
    [0, 1, 0, 0, 0, 0, 1, 0],
    [0, 0, 1, 1, 1, 1, 0, 0],
]

for y, row in enumerate(smiley):
    for x, pixel in enumerate(row):
        color = "#000000" if pixel else "#FFFFFF"
        set_pixel(x + 5, y + 5, color)
