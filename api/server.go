package api

import (
	"github.com/rs/zerolog"
)

// Server is an API server.
type Server struct {
	storage Storage
	log     zerolog.Logger
}

// NewServer will create a new API server.
func NewServer(storage Storage, log zerolog.Logger) *Server {

	server := Server{
		storage: storage,
		log:     log,
	}

	return &server
}
