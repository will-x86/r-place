curl --request POST \
  --url http://localhost:8081/api/pixels \
  --header 'Accept: */*' \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: curl/7.68.0' \
  --data '{"x":1, "y":2, "hex":"#102932"}'