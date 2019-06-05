package http

import "net/http"

// MethodNotAllowed replies to the request with an HTTP 405 not allowed error.
func MethodNotAllowed(w http.ResponseWriter, meth string) {
	http.Error(w, "("+meth+") "+http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}
