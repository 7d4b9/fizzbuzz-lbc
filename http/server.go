// Package http wrapps the underlying controller call and provide the from an http application standard endpoints.
package http

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Controller abstracts the fizzbuzz controller
type Controller interface {
	// Controller computes a fizzbuzz
	// all multiples of int1 are replaced by str1,
	// all multiples of int2 are replaced by str2,
	// all multiples of int1 and int2 are replaced by str1str2.
	FizzBuzz(int1, int2, limit int, str1, str2 string) (string, error)
}

// Handler computes the fissbuzz results
type Handler struct {
	Controller
}

// NewServer allocates and returns a new ServeMux.
func NewServer(controller Controller) *http.Server {
	mux := NewServeMux()
	handler := &Handler{
		controller,
	}
	mux.HandleFunc("/fizzbuzz", handler.FizzBuzz)
	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

// FizzBuzz handles an HTTP - FIZZBUZZ call and wraps the underlying application controller
func (app *Handler) FizzBuzz(w http.ResponseWriter, r *http.Request) {
	switch meth := r.Method; meth {
	case "POST":
		var bodyParameters struct {
			Int1  int    `json:"int1"`
			Int2  int    `json:"int2"`
			Limit int    `json:"limit"`
			Str1  string `json:"str1"`
			Str2  string `json:"str2"`
		}
		if err := json.NewDecoder(r.Body).Decode(&bodyParameters); err != nil {
			log.WithError(err).Error("missing request body parameters")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result, err := app.Controller.FizzBuzz(
			bodyParameters.Int1,
			bodyParameters.Int2,
			bodyParameters.Limit,
			bodyParameters.Str1,
			bodyParameters.Str2)
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"params": bodyParameters,
			}).Error("no fizzbuzz response payload")
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.Write([]byte(result))
	default:
		MethodNotAllowed(w, meth)
	}
}
