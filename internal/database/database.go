package database

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/lstratta/crosstech-submission/config"
	"github.com/lstratta/crosstech-submission/internal/models"
)

// func init() {
// 	// Register many to many model so ORM can better recognize m2m relation.
// 	// This should be done before dependant models are used.
// 	orm.RegisterTable((*models.TrackToSignal)(nil))
// }

type DB struct {
	conn *pg.DB
}

func (db *DB) Conn() *pg.DB {
	return db.conn
}

func SetupLocalDB(conf config.Config) (*DB, error) {
	opt, err := pg.ParseURL(conf.DatabaseURI)
	if err != nil {
		return nil, err
	}

	c := pg.Connect(opt)
	return &DB{conn: c}, nil
}

func MigrateModels(db *DB) error {

	for _, model := range defaultModels() {
		err := db.conn.Model(model).CreateTable(&orm.CreateTableOptions{IfNotExists: true})
		if err != nil {
			return err
		}
	}
	return nil
}

// Separate function to allow for easy addition of new
// models in future code updates
func defaultModels() []any {
	return []any{
		&models.TrackSignalJoin{},
		&models.Track{},
		&models.Signal{},
	}
}

func (db *DB) InsertData(ctx context.Context, track models.Track) {
	c := db.conn

	c.ExecContext(ctx, `
		INSERT INTO tracks (track_id, source, target)
		VALUES (?0, ?1, ?2);`,
		track.TrackId, track.Source, track.Target,
	)

	for _, s := range track.SignalIDs {
		c.ExecContext(ctx, `
			INSERT INTO signals (elr, mileage, signal_id, signal_name)
			VALUES (?0, ?1, ?2, ?3);`,
			s.ELR, s.Mileage, s.SignalId, s.SignalName,
		)

		c.ExecContext(ctx, `
			INSERT INTO track_signal_joins (signal_id, track_id)
			VALUES (?0, ?1);`,
			s.SignalId, track.TrackId,
		)
	}

}

func (db *DB) Tracks() ([]models.Track, error) {
	t := []models.Track{}
	_, err := db.conn.Query(&t, `
	  SELECT * FROM tracks;
	`)
	if err != nil {
		return nil, fmt.Errorf("error querying database for tracks: %v", err)
	}

	return t, nil
}

func (db *DB) Signals() ([]models.Signal, error) {
	s := []models.Signal{}
	_, err := db.conn.Query(&s, `
	  SELECT * FROM signals;
	`)
	if err != nil {
		return nil, fmt.Errorf("error querying database for tracks: %v", err)
	}

	return s, nil
}

func (db *DB) SignalsById(id string) ([]models.Signal, error) {
	s := []models.Signal{}
	_, err := db.conn.Query(&s, `
	  SELECT * FROM signals
	  WHERE signal_id = ?0;
	`, id)
	if err != nil {
		return nil, fmt.Errorf("error querying database for tracks: %v", err)
	}

	return s, nil
}

func (db *DB) TracksBySignalId(id string) ([]models.Track, error) {
	t := []models.Track{}

	return t, nil
}
