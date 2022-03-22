package api

// FIXME: Change this once the interface becomes clear.
type Storage interface {
	GetEvents(typ EventType, filter Filter) (interface{}, error)
}
