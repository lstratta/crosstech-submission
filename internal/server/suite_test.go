package server

import (
	"testing"

	"github.com/go-pg/pg/v10/orm"
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
	if err = database.MigrateModels(db); err != nil {
		ts.T().Errorf("error setting up database: %s", err)
	}
}

func (ts *TestSuite) SetupTest() {
	for _, t := range models.SetupOneTrackWithFiveSignals() {
		_, err := ts.srv.db.CreateTrack(&t)
		if err != nil {
			ts.T().Errorf("error creating tracks: %s", err)
		}
		for _, s := range t.SignalIds {
			_, err = ts.srv.db.CreateSignal(&s)
			if err != nil {
				ts.T().Errorf("error creating signals: %s", err)
			}
			err = ts.srv.db.CreateTrackSignalJoin(s.SignalId, t.TrackId)
			if err != nil {
				ts.T().Errorf("error creating track signal join table: %s", err)
			}
		}
	}
}

func (ts *TestSuite) AfterTest() {
	for _, m := range database.DefaultModels() {
		err := ts.srv.db.Conn().Model(m).DropTable(&orm.DropTableOptions{})
		if err != nil {
			ts.T().Errorf("error creating track signal join table: %s", err)
		}
	}
}

func TestSuite_Server(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
