package api

// Server is an API server.
type Server struct {
	storage Storage
}

// NewServer will create a new API server.
func NewServer(storage Storage) *Server {

	server := Server{
		storage: storage,
	}

	return &server
}
