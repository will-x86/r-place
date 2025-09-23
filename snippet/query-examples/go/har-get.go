package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://will-x86.com/api/pixelsq?x=1&y=2&hex=%23102932"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "curl/7.68.0")
	req.Header.Add("Accept", "*/*")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
