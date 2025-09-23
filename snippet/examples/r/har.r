library(httr)

url <- "https://will-x86.com/api/pixels"

payload <- "{\"x\":1, \"y\":2, \"hex\":\"#102932\"}"

encode <- "json"

response <- VERB("POST", url, body = payload, add_headers(User_Agent = 'curl/7.68.0', '), content_type("application/json"), accept("*/*"), encode = encode)

content(response, "text")
