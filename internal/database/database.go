// database configuration file
// the models that are used in the application are
// added here

package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
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
