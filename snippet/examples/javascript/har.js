const data = JSON.stringify({
  "x": 1,
  "y": 2,
  "hex": "#102932"
});

const xhr = new XMLHttpRequest();
xhr.withCredentials = true;

xhr.addEventListener("readystatechange", function () {
  if (this.readyState === this.DONE) {
    console.log(this.responseText);
  }
});

xhr.open("POST", "https://place.willx86.com/api/pixels");
xhr.setRequestHeader("Content-Type", "application/json");
xhr.setRequestHeader("User-Agent", "curl/7.68.0");
xhr.setRequestHeader("Accept", "*/*");

xhr.send(data);
