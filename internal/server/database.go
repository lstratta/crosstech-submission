package server

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/lstratta/crosstech-submission/internal/models"
)

func (s *Server) setupDB() error {
	c := s.conf
	opt, err := pg.ParseURL(c.DatabaseURI)
	if err != nil {
		return err
	}

	s.db = pg.Connect(opt)

	return nil
}

// Migrates all models and creates tables based on
// their attributes
func (s *Server) migrateModels() error {

	for _, model := range defaultModels() {
		err := s.db.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

// Separate function to allow for easy addition of new
// models in future code updates - shown for potential.
func defaultModels() []any {
	return []any{
		&models.TrackToSignal{}, // many-to-many table must come first
		&models.Track{},
		&models.Signal{},
	}
}
