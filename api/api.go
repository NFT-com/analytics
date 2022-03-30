package api

import (
	"github.com/rs/zerolog"
)

// FIXME: What will be the input format for the parameters - e.g. chain or collection? Our IDs?
// FIXME: What will be the input for the NFT then - also the ID?
// FIXME: What are the start/end times - timestamps or dates or?

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
