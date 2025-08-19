var client = new RestClient("http://localhost:8081/api/pixels");
var request = new RestRequest("", Method.Post);
request.AddHeader("Content-Type", "application/json");
request.AddHeader("User-Agent", "curl/7.68.0");
request.AddHeader("Accept", "*/*");
request.AddParameter("application/json", "{\"x\":1, \"y\":2, \"hex\":\"#102932\"}", ParameterType.RequestBody);
var response = client.Execute(request);