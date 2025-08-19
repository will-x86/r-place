const http = require('http');

const options = {
  method: 'POST',
  hostname: 'localhost',
  port: '8081',
  path: '/api/pixels',
  headers: {
    'Content-Type': 'application/json',
    'User-Agent': 'curl/7.68.0',
    Accept: '*/*'
  }
};

const req = http.request(options, function (res) {
  const chunks = [];

  res.on('data', function (chunk) {
    chunks.push(chunk);
  });

  res.on('end', function () {
    const body = Buffer.concat(chunks);
    console.log(body.toString());
  });
});

req.write(JSON.stringify({x: 1, y: 2, hex: '#102932'}));
req.end();