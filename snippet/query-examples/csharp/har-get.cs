var client = new RestClient("https://will-x86.com/api/pixelsq?x=1&y=2&hex=%23102932");
var request = new RestRequest(Method.GET);
request.AddHeader("User-Agent", "curl/7.68.0");
request.AddHeader("Accept", "*/*");
IRestResponse response = client.Execute(request);
