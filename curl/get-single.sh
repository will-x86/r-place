#!/bin/sh
printf "Via query params:(X-1,Y-1)\n"
curl -X POST "localhost:8081/api/pixel?x=1&y=1"
printf "\n\n"


printf "Via json body(X-1,Y-2):\n"
curl -X POST \
-H "Content-Type: application/json" \
-d '{"x":1,"y":2}' \
localhost:8081/api/pixel
printf "\n"
