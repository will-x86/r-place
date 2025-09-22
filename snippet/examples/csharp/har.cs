var client = new RestClient("https://place.willx86.com/api/pixels");
var request = new RestRequest(Method.POST);
request.AddHeader("Content-Type", "application/json");
request.AddHeader("User-Agent", "curl/7.68.0");
request.AddHeader("Accept", "*/*");
request.AddParameter("application/json", "{\"x\":1, \"y\":2, \"hex\":\"#102932\"}", ParameterType.RequestBody);
IRestResponse response = client.Execute(request);
