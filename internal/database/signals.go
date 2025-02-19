package database

import (
	"fmt"

	"github.com/lstratta/crosstech-submission/internal/models"
)

func (db *DB) Signals() ([]models.Signal, error) {
	s := []models.Signal{}
	_, err := db.conn.Query(&s, `
	  SELECT * FROM signals;
	`)
	if err != nil {
		return nil, fmt.Errorf("error querying database for signals: %v", err)
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
		return nil, fmt.Errorf("error querying database for signals: %v", err)
	}

	return s, nil
}
