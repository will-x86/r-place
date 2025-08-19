#!/bin/sh

if [ -z "$1" ]; then
  echo "Usage: ./set-multiple.sh <image_file>"
  exit 1
fi

API_URL="http://localhost:8081/api/pixels"
IMAGE_FILE="$1"
SIZE="1000x1000"

convert "$IMAGE_FILE" -resize $SIZE\! txt:- | \
sed -n 's/^\([0-9]*\),\([0-9]*\):.*\(#[0-9a-fA-F]\{6\}\).*/\1 \2 \3/p' | \
while read -r x y hex; do
  curl -s -X POST \
    -H "Content-Type: application/json" \
    -d "{\"x\":$x, \"y\":$y, \"hex\":\"$hex\"}" \
    "$API_URL" > /dev/null &
done

wait
