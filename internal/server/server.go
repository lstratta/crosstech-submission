package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lstratta/crosstech-submission/internal/config"
)

type Server struct {
	srv    *http.Server
	router *echo.Echo
}

func New(c config.Config) (*Server, error) {
	e := echo.New()

	return &Server{
		srv: &http.Server{
			Addr:    ":7777",
			Handler: e,
		},
		router: e,
	}, nil
}

func (s *Server) ListenAndServe() error {
	s.Routes()
	return s.srv.ListenAndServe()
}
