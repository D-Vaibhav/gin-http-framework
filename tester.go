package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://localhost:7000/videos"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	// test:test
	req.Header.Add("Authorization", "Basic dGVzdDp0ZXN0")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
