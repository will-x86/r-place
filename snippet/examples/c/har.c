CURL *hnd = curl_easy_init();

curl_easy_setopt(hnd, CURLOPT_CUSTOMREQUEST, "POST");
curl_easy_setopt(hnd, CURLOPT_URL, "https://place.willx86.com/api/pixels");

struct curl_slist *headers = NULL;
headers = curl_slist_append(headers, "Content-Type: application/json");
headers = curl_slist_append(headers, "User-Agent: curl/7.68.0");
headers = curl_slist_append(headers, "Accept: */*");
curl_easy_setopt(hnd, CURLOPT_HTTPHEADER, headers);

curl_easy_setopt(hnd, CURLOPT_POSTFIELDS, "{\"x\":1, \"y\":2, \"hex\":\"#102932\"}");

CURLcode ret = curl_easy_perform(hnd);
