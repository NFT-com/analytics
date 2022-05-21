package api

import (
	"github.com/rs/zerolog"
)

// Server is an API server.
type Server struct {
	log            zerolog.Logger
	storage        Storage
	aggregationAPI Stats

	searchLimit uint
}

// NewServer creates a new API server.
func NewServer(log zerolog.Logger, storage Storage, aggregationAPI Stats, opts ...OptionFunc) *Server {

	cfg := defaultConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	server := Server{
		log:            log,
		storage:        storage,
		aggregationAPI: aggregationAPI,
		searchLimit:    cfg.SearchLimit,
	}

	return &server
}
