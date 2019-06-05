package http

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
	// "githib.com/sirups
)

// Controller abstracts the fizzbuzz controller
type Controller interface {
	// Controller computes a fizzbuzz request
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
	handler := &Handler{
		controller,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.FizzBuzz)
	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

// FizzBuzz handles an HTTP - FIZZBUZZ call and wraps the underlying application engine
func (app *Handler) FizzBuzz(w http.ResponseWriter, r *http.Request) {
	switch meth := r.Method; meth {
	case "POST":
		// Decode the JSON in the body and overwrite 'tom' with it
		d := json.NewDecoder(r.Body)
		var p struct {
			Int1  int    `json:"int1"`
			Int2  int    `json:"int2"`
			Limit int    `json:"limit"`
			Str1  string `json:"str1"`
			Str2  string `json:"str2"`
		}
		if err := d.Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		text, err := app.Controller.FizzBuzz(p.Int1, p.Int2, p.Limit, p.Str1, p.Str2)
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"params": p,
			}).Error("fizzbuzz failure")
		}
		w.Write([]byte(text))
	default:
		MethodNotAllowed(w, meth)
	}
}
