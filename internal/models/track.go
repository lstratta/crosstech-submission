package models

type Track struct {
	TrackID   int
	Source    string
	Target    string
	SignalIDs []SignalIDs
}

type SignalIDs struct {
	SignalID   int
	SignalName string
	ELR        string
	Mileage    float32
}
