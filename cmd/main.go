package main

import (
	"log"
	"net/http"

	"github.com/lstratta/crosstech-submission/config"
	"github.com/lstratta/crosstech-submission/internal/server"
)

// generate the config using environment variables,
// create the server, and run the server
func main() {
	conf := config.New()

	s, err := server.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
