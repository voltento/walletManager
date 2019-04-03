package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

func sendRequest(url string, method string, body string) (string, error) {
	var jsonStr = []byte(body)
	req, er := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if er != nil {
		return "", er
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if data, er := ioutil.ReadAll(resp.Body); er != nil {
		return "", er
	} else {
		return strings.Trim(string(data), "\n"), er
	}
}
