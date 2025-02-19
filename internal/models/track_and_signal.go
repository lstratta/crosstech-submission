package models

type Track struct {
	TrackPk   int64    `json:"-" pg:",pk"`
	TrackId   int      `json:"track_id" pg:",unique"`
	Source    string   `json:"source"`
	Target    string   `json:"target"`
	SignalIds []Signal `json:"signal_ids,omitempty"`
}

type Signal struct {
	SignalPk   int64   `json:"-" pg:",pk"`
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
	Tracks  []Track  `json:"tracks"`
	Message string   `json:"message,omitempty"`
	Errors  []Errors `json:"errors,omitempty"`
}

type SignalResponse struct {
	Signals []Signal `json:"signals"`
	Message string   `json:"message,omitempty"`
	Errors  []Errors `json:"errors,omitempty"`
}

type Errors struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}
