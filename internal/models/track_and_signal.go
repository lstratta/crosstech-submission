package models

type Track struct {
	TrackPK   int64    `json:"track_pk" pg:",pk"`
	TrackId   int      `json:"track_id"`
	Source    string   `json:"source"`
	Target    string   `json:"target"`
	SignalIDs []Signal `json:"signal_ids,omitempty"`
}

type Signal struct {
	SignalPK   int64   `json:"signal_pk" pg:",pk"`
	SignalId   int     `json:"signal_id"`
	SignalName string  `json:"signal_name"`
	ELR        string  `json:"elr"`
	Mileage    float32 `json:"mileage"`
}

type TrackSignalJoin struct {
	Id       int64 `pg:",pk"`
	SignalId int64
	TrackId  int64
}

type TrackResponse struct {
	Tracks []Track `json:"tracks"`
}

type SignalResponse struct {
	Signals []Signal `json:"signals"`
}
