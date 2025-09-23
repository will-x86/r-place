const data = null;

const xhr = new XMLHttpRequest();
xhr.withCredentials = true;

xhr.addEventListener("readystatechange", function () {
  if (this.readyState === this.DONE) {
    console.log(this.responseText);
  }
});

xhr.open("GET", "https://will-x86.com/api/pixelsq?x=1&y=2&hex=%23102932");
xhr.setRequestHeader("User-Agent", "curl/7.68.0");
xhr.setRequestHeader("Accept", "*/*");

xhr.send(data);
