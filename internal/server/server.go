// server.go handles all the server creation code
// this allows us to make a robust server with a lot of
// extensible features
package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lstratta/crosstech-submission/config"
	"github.com/lstratta/crosstech-submission/internal/data"
	"github.com/lstratta/crosstech-submission/internal/database"
	"github.com/lstratta/crosstech-submission/internal/models"
)

// the server is based around this struct, which uses the echo
// framework as a router and http.handler
type Server struct {
	srv    *http.Server
	router *echo.Echo
	db     *database.DB
	conf   config.Config
}

func New(c config.Config) (*Server, error) {
	e := echo.New()

	// instantiate the server
	s := &Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", c.Host, c.Port),
			Handler: e,
		},
		router: e,
		conf:   c,
	}

	// setup a local db connection
	db, err := database.SetupLocalDB(s.conf)
	if err != nil {
		return nil, fmt.Errorf("error setting up database: %v", err)
	}

	s.db = db

	// check db connectivity
	ctx := context.Background()
	if err := s.db.Conn().Ping(ctx); err != nil {
		return nil, fmt.Errorf("error pinging database - it may not be ready: %v", err)
	}

	// create tables
	if err := database.MigrateTables(s.db); err != nil {
		return nil, fmt.Errorf("error migrating tables: %v", err)
	}

	// populate database
	if err := hydrateDb(s); err != nil {
		return nil, fmt.Errorf("error hydrating db: %v", err)
	}

	return s, nil
}

// the function that implements the middleware, routes,
// and starts the server
func (s *Server) ListenAndServe() error {
	s.middleware()
	s.routes()
	return s.srv.ListenAndServe()
}

func hydrateDb(s *Server) error {
	// insert json data
	trackData, err := data.ParseJsonData()
	if err != nil {
		return fmt.Errorf("error inputing data: %v", err)
	}

	// prevent duplication of data on startup and reload
	res, err := s.db.Conn().Query(&models.Track{}, "SELECT * FROM tracks;")
	if res.RowsReturned() < 1 {
		fmt.Println("no data found in db... populating...")
		for _, t := range trackData {
			s.db.CreateTrackWithSignals(&t)
		}
	} else if err != nil {
		return fmt.Errorf("error querying db: %v", err)
	}

	return nil
}
