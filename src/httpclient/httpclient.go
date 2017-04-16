package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func My_Http_Do() {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:9090", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=annay")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	resp, err = http.Head("http://localhost:9090")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
