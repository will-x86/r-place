#!/bin/sh
if [ -z "$1" ] || [ -z "$2" ] || [ -z "$3" ]; then
  echo "Usage: ./set-image.sh <image_file> <x> <y>"
  exit 1
fi

IMAGE_FILE="$1"
X="$2"
Y="$3"
API_URL="https://will-x86.com/api/image10x10"

convert "$IMAGE_FILE" -resize 10x10\! "$IMAGE_FILE.tmp.png"

curl -X POST -F "x=$X" -F "y=$Y" -F "image=@$IMAGE_FILE.tmp.png" "$API_URL"

rm "$IMAGE_FILE.tmp.png"
