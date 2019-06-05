package fizzbuzz

import "testing"

func TestControllerFizzBuzz(t *testing.T) {
	var app Controller
	const expectedText = "1 fizz buzz fizz 5 fizzbuzz 7 fizz buzz fizz"
	text, err := app.FizzBuzz(2, 3, 10, "fizz", "buzz")
	if err != nil {
		t.Error("fizzbuzz error,", err)
	}
	if text != expectedText {
		t.Error("unexpected text, got '", text, "' instead of expected '", expectedText)
	}
}
