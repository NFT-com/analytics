package api

import (
	"github.com/rs/zerolog"
)

// Server is an API server.
type Server struct {
	storage Storage
	log     zerolog.Logger

	searchLimit uint
}

// NewServer creates a new API server.
func NewServer(storage Storage, log zerolog.Logger, opts ...OptionFunc) *Server {

	cfg := defaultConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	server := Server{
		storage:     storage,
		log:         log,
		searchLimit: cfg.SearchLimit,
	}

	return &server
}
