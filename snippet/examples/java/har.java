HttpResponse<String> response = Unirest.post("https://place.willx86.com/api/pixels")
  .header("Content-Type", "application/json")
  .header("User-Agent", "curl/7.68.0")
  .header("Accept", "*/*")
  .body("{\"x\":1, \"y\":2, \"hex\":\"#102932\"}")
  .asString();
