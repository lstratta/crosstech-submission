package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/lstratta/crosstech-submission/internal/config"
)

type Server struct {
	srv    *http.Server
	router *echo.Echo
	db     *pg.DB
	conf   config.Config
}

func New(c config.Config) (*Server, error) {
	e := echo.New()

	s := &Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", c.Host, c.Port),
			Handler: e,
		},
		router: e,
		conf:   c,
	}

	if err := s.setupDB(); err != nil {
		return nil, err
	}

	ctx := context.Background()

	if err := s.db.Ping(ctx); err != nil {
		panic(err)
	}

	s.migrateModels()

	return s, nil
}

func (s *Server) ListenAndServe() error {
	s.middleware(s.conf)
	s.routes()
	return s.srv.ListenAndServe()
}
