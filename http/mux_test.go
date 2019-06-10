package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMuxStatistcsHandler(t *testing.T) {
	mux := NewServeMux()

	mux.HandleFunc("/api/foo", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/api/bar", func(w http.ResponseWriter, r *http.Request) {})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	http.Get(ts.URL + "/api/bar")

	resp, err := http.Get(ts.URL + "/statistics")
	if err != nil {
		t.Error(err)
	}
	expectedBody := "{\"Pattern\":\"/api/bar\",\"Count\":1}"
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("cannot read response body,", err)
	}
	if strings.TrimSpace(string(body)) != expectedBody {
		t.Error("unexpected body, got", string(body), ", instead of expected:", expectedBody)
	}

	http.Get(ts.URL + "/api/foo")
	http.Get(ts.URL + "/api/foo")

	resp, err = http.Get(ts.URL + "/statistics")
	if err != nil {
		t.Error(err)
	}
	expectedBody = "{\"Pattern\":\"/api/foo\",\"Count\":2}"
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("cannot read response body,", err)
	}
	if strings.TrimSpace(string(body)) != expectedBody {
		t.Error("unexpected body, got", string(body), ", instead of expected:", expectedBody)
	}
}
