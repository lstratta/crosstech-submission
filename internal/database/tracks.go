// tracks.go file holds all the functions relating to database access for tracks
// This helps to reduce file size and makes it easier to maintain
package database

import (
	"fmt"

	"github.com/lstratta/crosstech-submission/internal/models"
)

func (db *DB) Tracks() ([]models.Track, error) {
	ts := []models.Track{}
	_, err := db.conn.Query(&ts, `
	  SELECT * 
	  FROM tracks
	  ORDER BY track_id ASC;
	`)
	if err != nil {
		return nil, fmt.Errorf("error querying database for tracks: %v", err)
	}

	return ts, nil
}

func (db *DB) TracksById(id int) ([]models.Track, error) {
	ts := []models.Track{}
	_, err := db.conn.Query(&ts, `
	  SELECT * FROM tracks
	  WHERE track_id = ?0;
	`, id)
	if err != nil {
		return nil, fmt.Errorf("error querying database for tracks: %v", err)
	}

	// add signals to corresponding track
	for n, t := range ts {
		s := []models.Signal{}

		_, err := db.conn.Query(&s, `
	  	SELECT signals.signal_id, signals.signal_name, signals.elr, signals.mileage
		FROM tracks
		JOIN track_signal_joins ON tracks.track_id = track_signal_joins.track_id
		JOIN signals ON track_signal_joins.signal_id = signals.signal_id
		WHERE tracks.track_id = ?0;
	`, t.TrackId)
		if err != nil {
			return nil, fmt.Errorf("error querying database for tracks: %v", err)
		}

		ts[n].SignalIds = s
	}

	return ts, nil
}

func (db *DB) TracksBySignalId(id int) ([]models.Track, error) {
	t := []models.Track{}

	_, err := db.conn.Query(&t, `
	  	SELECT DISTINCT tracks.track_id, tracks.source, tracks.target
		FROM signals
		JOIN track_signal_joins ON signals.signal_id = track_signal_joins.signal_id
		JOIN tracks ON track_signal_joins.track_id = tracks.track_id
		WHERE signals.signal_id = ?0;
	`, id)
	if err != nil {
		return nil, fmt.Errorf("error querying database for tracks: %v", err)
	}

	return t, nil
}

func (db *DB) CreateTrackWithSignals(t *models.Track) (*models.Track, error) {
	_, err := db.conn.Exec(`
		INSERT INTO tracks (track_id, source, target)
		VALUES (?0, ?1, ?2)
		ON CONFLICT (track_id)
		DO UPDATE SET track_id = ?0, source = ?1, target = ?2;
	`, t.TrackId, t.Source, t.Target)
	if err != nil {
		return nil, fmt.Errorf("error inserting track into database: %v", err)
	}

	if len(t.SignalIds) < 1 {
		return t, nil
	}

	for _, s := range t.SignalIds {
		_, err := db.CreateSignal(&s)
		if err != nil {
			return nil, err
		}

		err = db.CreateTrackSignalJoin(s.SignalId, t.TrackId)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

func (db *DB) CreateTrackSignalJoin(sId, tId int) error {
	_, err := db.conn.Exec(`
			INSERT INTO track_signal_joins (signal_id, track_id)
			VALUES (?0, ?1);`,
		sId, tId)
	if err != nil {
		return fmt.Errorf("error inserting tracks_signal_join into database: %v", err)
	}

	return nil
}

func (db *DB) UpdateTrack(t *models.Track) (*models.Track, error) {
	_, err := db.conn.Exec(
		`
			UPDATE tracks
			SET source = ?0,
				target = ?1
			WHERE track_id = ?2;
		`, t.Source, t.Target, t.TrackId)
	if err != nil {
		return nil, fmt.Errorf("error updating track: %v", err)
	}

	return t, nil
}

func (db *DB) DeleteTrackById(id int) error {
	_, err := db.conn.Exec(
		`
			DELETE FROM tracks
			WHERE track_id = ?0;
		`, id)
	if err != nil {
		return fmt.Errorf("error deleting track: %v", err)
	}

	return nil
}
