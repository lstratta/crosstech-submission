package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/lstratta/crosstech-submission/config"
	"github.com/lstratta/crosstech-submission/internal/data"
	"github.com/lstratta/crosstech-submission/internal/database"
	"github.com/lstratta/crosstech-submission/internal/models"
)

type Server struct {
	srv    *http.Server
	router *echo.Echo
	db     *database.DB
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

	db, err := database.SetupLocalDB(s.conf)
	if err != nil {
		return nil, fmt.Errorf("error setting up database: %v", err)
	}

	s.db = db

	ctx := context.Background()
	if err := s.db.Conn().Ping(ctx); err != nil {
		return nil, fmt.Errorf("error pinging database - it may not be ready: %v", err)
	}

	if err := database.MigrateModels(s.db); err != nil {
		return nil, fmt.Errorf("error migrating tables: %v", err)
	}

	trackData, err := data.ParseJsonData()
	if err != nil {
		return nil, fmt.Errorf("error inputing data: %v", err)
	}

	// prevent duplication of data on startup and reload
	res, err := s.db.Conn().Query(&models.Track{}, "SELECT * FROM tracks;")
	if pg.Result.RowsReturned(res) == 0 {
		fmt.Println("no data found in db... populating...")
		for _, t := range trackData {
			s.db.InsertData(ctx, t)
		}
	} else if err != nil {
		return nil, fmt.Errorf("error querying db: %v", err)
	}

	return s, nil
}

func (s *Server) ListenAndServe() error {
	s.middleware()
	s.routes()
	return s.srv.ListenAndServe()
}
