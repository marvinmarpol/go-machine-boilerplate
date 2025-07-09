package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

var netTransport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 10 * time.Second,
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 100,
}
var netClient = &http.Client{
	Timeout:   30 * time.Second,
	Transport: netTransport,
}

func Request(request *http.Request) (*http.Response, error) {
	netClient.Timeout = 30 * time.Second
	resp, err := netClient.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func requestWithTimeout(request *http.Request, timeoutInSecond int) (*http.Response, error) {
	netClient.Timeout = time.Duration(timeoutInSecond) * time.Second
	resp, err := netClient.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateRequestAndDo func to create *http.Request and return *http.Response
func CreateRequestAndDo(url, method string, timeout int, headers map[string]string, requestPayload interface{}) (data *http.Response, err error) {
	// marshal payload to json
	jsonData, err := json.Marshal(requestPayload)
	if err != nil {
		return nil, err
	}
	// create request object
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if err != nil {
		return nil, err
	}

	// do the http request
	data, err = requestWithTimeout(req, timeout)
	return
}

// CreateRequestAndParse func to create *http.Request and return the payload response
func CreateRequestAndParse(url, method string, timeout int, headers map[string]string, requestPayload, returnPointer interface{}) (data *http.Response, err error) {
	response, err := CreateRequestAndDo(url, method, timeout, headers, requestPayload)
	if err != nil {
		return nil, err
	}

	return response, ParseResponseBody(response, returnPointer)
}

// ParseResponseBody is a func to parse json response body to map of string to interface
func ParseResponseBody(response *http.Response, returnPointer interface{}) (err error) {
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if len(responseData) > 0 {
		err = json.Unmarshal(responseData, &returnPointer)
	}
	return
}

// GetRequestRawQuery is a function to get http request raw query (example: '?phone=+6281xxx&ktp=32751xxx')
func GetRequestRawQuery(request *http.Request) string {
	url, err := url.Parse(request.RequestURI)
	if err != nil {
		return ""
	}

	return url.RawQuery
}
