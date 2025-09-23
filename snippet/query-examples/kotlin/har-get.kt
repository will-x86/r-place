val client = OkHttpClient()

val request = Request.Builder()
  .url("https://will-x86.com/api/pixelsq?x=1&y=2&hex=%23102932")
  .get()
  .addHeader("User-Agent", "curl/7.68.0")
  .addHeader("Accept", "*/*")
  .build()

val response = client.newCall(request).execute()
