const data = JSON.stringify({
  x: 1,
  y: 2,
  hex: '#102932'
});

const xhr = new XMLHttpRequest();
xhr.withCredentials = true;

xhr.addEventListener('readystatechange', function () {
  if (this.readyState === this.DONE) {
    console.log(this.responseText);
  }
});

xhr.open('POST', 'http://localhost:8081/api/pixels');
xhr.setRequestHeader('Content-Type', 'application/json');
xhr.setRequestHeader('User-Agent', 'curl/7.68.0');
xhr.setRequestHeader('Accept', '*/*');

xhr.send(data);