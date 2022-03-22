package api

// API provides the Events REST API functionality.
type API struct {
	storage Storage
}

// New creates a new API handler.
func New(storage Storage) *API {

	api := API{
		storage: storage,
	}

	return &api
}
