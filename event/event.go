package event

//Event simple event
type Event struct {
	id Type
}

//Type enum for eventIDs
type Type int

const (
	TestEvent Type = iota
)
