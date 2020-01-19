package ipipgo

import "net/http"

var client = http.DefaultClient

func SetClient(c *http.Client) {
	client = c
}

var header = http.Header{}

func SetHeader(h http.Header) {
	header = h
}

func httpGet(url string) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header = header
	return client.Do(req)
}
