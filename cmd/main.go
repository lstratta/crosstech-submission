package main

import (
	"log"
	"net/http"

	"github.com/lstratta/crosstech-submission/internal/server"
)

func main() {
	s := server.New()

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
