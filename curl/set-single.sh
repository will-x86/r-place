#!/bin/sh
printf "Setting X-1,Y-2 to #102932\n"
curl -X POST \
-H "Content-Type: application/json" \
-d '{"x":1, "y":2, "hex":"#102932"}' \
https://will-x86.com/api/pixels


