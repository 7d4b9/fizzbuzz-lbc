package http

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

type DummyController struct{}

func (d *DummyController) FizzBuzz(int1, int2, limit int, str1, str2 string) (string, error) {
	return "1 fizz buzz fizz 5 fizzbuzz 7 fizz buzz fizz", nil
}

func TestFizzBuzzHandler(t *testing.T) {
	handler := (&Handler{&DummyController{}}).FizzBuzz
	req := httptest.NewRequest("POST", "http://example.com/foo", bytes.NewBuffer([]byte(
		`{"Int1": 2,"Int2":3 , "Limit": 10, "Str1": "fizz", "Str2": "buzz"}`,
	)))
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()

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
