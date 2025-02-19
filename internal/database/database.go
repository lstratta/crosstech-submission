// database configuration file
// the models that are used in the application are
// added here

package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/lstratta/crosstech-submission/config"
	"github.com/lstratta/crosstech-submission/internal/models"
)

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

func MigrateTables(db *DB) error {

	// for range DefaultModels() {
	_, err := db.conn.Exec(`
		CREATE TABLE IF NOT EXISTS track_signal_joins (
			id SERIAL PRIMARY KEY,
			track_id BIGINT,
			signal_id BIGINT
			);`)

	if err != nil {
		return err
	}

	_, err = db.conn.Exec(`
		CREATE TABLE IF NOT EXISTS tracks (
			track_pk SERIAL PRIMARY KEY,
			track_id INT UNIQUE NOT NULL,
			source VARCHAR(100),
			target VARCHAR(100)
			);`)
	if err != nil {
		return err
	}

	_, err = db.conn.Exec(`
		CREATE TABLE IF NOT EXISTS signals (
			signal_pk SERIAL PRIMARY KEY,
			signal_id BIGINT NOT NULL,
			signal_name VARCHAR(100),
			elr VARCHAR(100),
			mileage REAL
			);`)
	if err != nil {
		return err
	}

	return nil
}

// Separate function to allow for easy addition of new
// models in future code updates
func DefaultModels() []any {
	return []any{
		&models.TrackSignalJoin{},
		&models.Track{},
		&models.Signal{},
	}
}
