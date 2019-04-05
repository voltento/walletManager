package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

type httpResp struct {
	data string
	code int
}

func sendRequest(url string, method string, body string) (httpResp, error) {
	var jsonStr = []byte(body)
	req, er := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if er != nil {
		return httpResp{}, er
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if data, er := ioutil.ReadAll(resp.Body); er != nil {
		return httpResp{}, er
	} else {
		return httpResp{data: strings.Trim(string(data), "\n"), code: resp.StatusCode}, er
	}
}
