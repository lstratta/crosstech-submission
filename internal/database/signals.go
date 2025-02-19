// signals.go file holds all the functions relating to database access for signals
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

	fmt.Println("DB SIGNALS, ", len(s))

	return s, nil
}

func (db *DB) SignalsById(id int) ([]models.Signal, error) {
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

func (db *DB) CreateSignal(s *models.Signal) (*models.Signal, error) {
	_, err := db.conn.Exec(`
			INSERT INTO signals (elr, mileage, signal_id, signal_name)
			VALUES (?0, ?1, ?2, ?3);`,
		s.ELR, s.Mileage, s.SignalId, s.SignalName)
	if err != nil {
		return nil, fmt.Errorf("error inserting signal into database: %v", err)
	}

	return s, nil
}

func (db *DB) UpdateSignal(s *models.Signal) (*models.Signal, error) {
	_, err := db.conn.Exec(
		`
			UPDATE signals
			SET signal_name = ?0,
				elr = ?1,
				mileage = ?2
			WHERE signal_id = ?3;
		`, s.SignalName, s.ELR, s.Mileage, s.SignalId)
	if err != nil {
		return nil, fmt.Errorf("error updating signal: %v", err)
	}

	return s, nil
}

func (db *DB) DeleteSignalById(id int) error {
	_, err := db.conn.Exec(
		`
			DELETE FROM signals
			WHERE signal_id = ?0;
		`, id)
	if err != nil {
		return fmt.Errorf("error deleting signal: %v", err)
	}

	return nil
}
