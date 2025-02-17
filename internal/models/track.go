package models

type Track struct {
	TrackId   int64 `pg:",pk"`
	Source    string
	Target    string
	SignalIDs []Signal `pg:"many2many:track_to_signals"`
}

type Signal struct {
	SignalId   int64 `pg:",pk"`
	SignalName string
	ELR        string
	Mileage    float32
}

// A table to track the many to many relations
// of tracks to signals
type TrackToSignal struct {
	TrackId  int
	SignalId int
}
