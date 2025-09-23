library(httr)

url <- "https://will-x86.com/api/pixelsq"

queryString <- list(
  x = "1",
  y = "2"
  hex = "#102932",
)

response <- VERB("GET", url, add_headers(User_Agent = 'curl/7.68.0', '), query = queryString, content_type("application/octet-stream"), accept("*/*"))

content(response, "text")
