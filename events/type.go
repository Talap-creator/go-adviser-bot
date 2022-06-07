package events

type fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type processor interface {
	Process(e Event) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type type
	Text string
}
