package api

import (
	"github.com/rs/zerolog"
)

// API provides the Events REST API functionality.
type API struct {
	storage Storage
	log     zerolog.Logger
}

// New creates a new API handler.
func New(storage Storage, log zerolog.Logger) *API {

	api := API{
		storage: storage,
		log:     log,
	}

	return &api
}
