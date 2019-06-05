// +build integration

package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHTTPServerFizzbuzz(t *testing.T) {
	reqBody := bytes.NewBuffer([]byte(
		`{"Int1": 2,"Int2":3 , "Limit": 10, "Str1": "fizz", "Str2": "buzz"}`,
	))
	req, err := http.NewRequest("POST", "http://fizzbuzz:8080", reqBody)
	if err != nil {
		t.Fatal("request error,", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("post error,", err)
	}

	const expectedStatusCode = 200
	var statusCode = resp.StatusCode
	if expectedStatusCode != statusCode {
		t.Error("unexpected status code, got '", statusCode, "' instead of expected '", expectedStatusCode)
	}

	const expectedContentType = "text/plain; charset=utf-8"
	var contentType = resp.Header.Get("Content-Type")
	if contentType != expectedContentType {
		t.Error("unexpected content type, got '", contentType, "' instead of expected '", expectedContentType)
	}

	const expectedBody = "1 fizz buzz fizz 5 fizzbuzz 7 fizz buzz fizz"
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("cannot read response body,", err)
	}
	if string(body) != expectedBody {
		t.Error("unexpected body, got '", string(body), "' instead of expected '", expectedBody)
	}

}
