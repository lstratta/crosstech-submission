package server

import (
	"testing"

	"github.com/lstratta/crosstech-submission/config"
	"github.com/lstratta/crosstech-submission/internal/database"
	"github.com/lstratta/crosstech-submission/internal/models"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	srv *Server
}

func (ts *TestSuite) SetupSuite() {
	conf := config.New()
	db, err := database.SetupLocalDB(conf)
	if err != nil {
		ts.T().Errorf("error setting up database: %s", err)
	}
	ts.srv = &Server{
		conf: conf,
		db:   db,
	}
}

func (ts *TestSuite) SetupTest() {
	// drop all tables
	_, err := ts.srv.db.Conn().Exec(`DROP TABLE IF EXISTS track_signal_joins;`)
	if err != nil {
		ts.T().Errorf("error dropping track signal join table: %s", err)
	}

	_, err = ts.srv.db.Conn().Exec(`DROP TABLE IF EXISTS tracks;`)
	if err != nil {
		ts.T().Errorf("error dropping track signal join table: %s", err)
	}

	_, err = ts.srv.db.Conn().Exec(`DROP TABLE IF EXISTS signals;`)
	if err != nil {
		ts.T().Errorf("error dropping track signal join table: %s", err)
	}

	// create the tables
	if err := database.MigrateTables(ts.srv.db); err != nil {
		ts.T().Errorf("error setting up database: %s", err)
	}

	// populate tables with test data
	for _, t := range models.SetupThreeTracksEachWithFiveSignals() {
		_, err := ts.srv.db.CreateTrackWithSignals(&t)
		if err != nil {
			ts.T().Errorf("error creating tracks: %s", err)
		}
	}
}

func TestSuite_Server(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
