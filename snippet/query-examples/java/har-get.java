HttpResponse<String> response = Unirest.get("https://will-x86.com/api/pixelsq?x=1&y=2&hex=%23102932")
  .header("User-Agent", "curl/7.68.0")
  .header("Accept", "*/*")
  .asString();
