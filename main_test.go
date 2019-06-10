// +build integration

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestHTTPServerFizzbuzz(t *testing.T) {
	reqBody := bytes.NewBuffer([]byte(
		`{"Int1": 2,"Int2":3 , "Limit": 10, "Str1": "fizz", "Str2": "buzz"}`,
	))
	req, err := http.NewRequest("POST", "http://fizzbuzz:8080/fizzbuzz", reqBody)
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

func TestHTTPServerStatistics(t *testing.T) {
	payload1 := []byte(`{"Int1": 2,"Int2":3 , "Limit": 10, "Str1": "fizz", "Str2": "buzz"}`)
	payload2 := []byte(`{"Int1": 2,"Int2":3 , "Limit": 10, "Str1": "easy", "Str2": "peasy"}`)
	for _, reqBody := range [][]byte{
		payload1,
		payload2,
		payload2,
		payload2,
	} {
		_, err := http.Post("http://fizzbuzz:8080/fizzbuzz", "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatal("post error,", err)
		}
	}
	resp, err := http.Get("http://fizzbuzz:8080/statistics")
	if err != nil {
		t.Fatal("post error,", err)
	}
	var stats json.RawMessage
	json.NewDecoder(resp.Body).Decode(&stats)
	expected := string([]byte(`{"Pattern":"/fizzbuzz/{\"Int1\": 2,\"Int2\":3 , \"Limit\": 10, \"Str1\": \"easy\", \"Str2\": \"peasy\"}","Count":3}`))
	got := string(stats)
	if got != expected {
		t.Error("unexpected body, got '", got, "' instead of expected '", expected)
	}
	strings.Trim(expected, `\`)
	http.Get("http://fizzbuzz:8080/statistics")
	http.Get("http://fizzbuzz:8080/statistics")
	http.Get("http://fizzbuzz:8080/statistics")
	resp, err = http.Get("http://fizzbuzz:8080/statistics")
	json.NewDecoder(resp.Body).Decode(&stats)
	expected = string([]byte(`"Pattern":"/statistics"`))
	got = string(stats)
	if !strings.Contains(got, expected) {
		t.Error("unexpected body, got '", got, "' does not contains expected '", expected)
	}
}
