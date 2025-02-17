package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	srv    *http.Server
	router *echo.Echo
}

func New() *Server {
	e := echo.New()

	return &Server{
		srv: &http.Server{
			Addr:    ":7777",
			Handler: e,
		},
		router: e,
	}
}

func (s *Server) ListenAndServe() error {
	s.Routes()
	return s.srv.ListenAndServe()
}
