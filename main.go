package main

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/bbrod/fizzbuzz-lbc/fizzbuzz"
	"gitlab.com/bbrod/fizzbuzz-lbc/http"
)

func main() {
	controller := &fizzbuzz.Controller{}
	server := http.NewServer(controller)
	if err := server.ListenAndServe(); err != nil {
		log.WithError(err).Error("HTTP server has left")
	}
}
