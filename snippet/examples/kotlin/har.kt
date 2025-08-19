val client = OkHttpClient()

val mediaType = MediaType.parse("application/json")
val body = RequestBody.create(mediaType, "{\"x\":1, \"y\":2, \"hex\":\"#102932\"}")
val request = Request.Builder()
  .url("http://localhost:8081/api/pixels")
  .post(body)
  .addHeader("Content-Type", "application/json")
  .addHeader("User-Agent", "curl/7.68.0")
  .addHeader("Accept", "*/*")
  .build()

val response = client.newCall(request).execute()