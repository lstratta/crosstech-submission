package models

type Track struct {
	TrackPK   int64    `pg:",pk"`
	TrackId   int      `json:"track_id"`
	Source    string   `json:"source"`
	Target    string   `json:"target"`
	SignalIDs []Signal `json:"signal_ids" pg:"has_many"`
}

type Signal struct {
	SignalPK   int64   `pg:",pk"`
	SignalId   int     `json:"signal_id"`
	SignalName string  `json:"signal_name"`
	ELR        string  `json:"elr"`
	Mileage    float32 `json:"mileage"`
}

// // A table to track the many to many relations
// // of tracks to signals
// type TrackToSignal struct {
// 	TrackToSignalPk int64 `pg:",pk"`
// 	TrackId         int
// 	SignalId        int
// }
