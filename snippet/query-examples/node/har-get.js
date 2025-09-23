const http = require("https");

const options = {
  "method": "GET",
  "hostname": "will-x86.com",
  "port": null,
  "path": "/api/pixelsq?x=1&y=2&hex=%23102932",
  "headers": {
    "User-Agent": "curl/7.68.0",
    "Accept": "*/*"
  }
};

const req = http.request(options, function (res) {
  const chunks = [];

  res.on("data", function (chunk) {
    chunks.push(chunk);
  });

  res.on("end", function () {
    const body = Buffer.concat(chunks);
    console.log(body.toString());
  });
});

req.end();
