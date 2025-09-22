#!/bin/sh
printf "Via query params:(X-1,Y-1)\n"
curl -X POST "https://place.willx86.com/api/pixel?x=1&y=1"
printf "\n\n"


printf "Via json body(X-1,Y-2):\n"
curl -X POST \
-H "Content-Type: application/json" \
-d '{"x":1,"y":2}' \
https://place.willx86.com/api/pixel
printf "\n"
