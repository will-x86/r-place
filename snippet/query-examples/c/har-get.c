CURL *hnd = curl_easy_init();

curl_easy_setopt(hnd, CURLOPT_CUSTOMREQUEST, "GET");
curl_easy_setopt(hnd, CURLOPT_URL, "https://will-x86.com/api/pixelsq?x=1&y=2&hex=%23102932");

struct curl_slist *headers = NULL;
headers = curl_slist_append(headers, "User-Agent: curl/7.68.0");
headers = curl_slist_append(headers, "Accept: */*");
curl_easy_setopt(hnd, CURLOPT_HTTPHEADER, headers);

CURLcode ret = curl_easy_perform(hnd);
