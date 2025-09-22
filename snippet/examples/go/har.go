package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://place.willx86.com/api/pixels"

	payload := strings.NewReader("{\"x\":1, \"y\":2, \"hex\":\"#102932\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "curl/7.68.0")
	req.Header.Add("Accept", "*/*")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
