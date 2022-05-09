package storage

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	server "github.com/NFT-com/graph-api/graph/api"
	"github.com/NFT-com/graph-api/graph/models/api"
)

// Network retrieves a single network based on the ID.
func (s *Storage) Network(id string) (*api.Network, error) {

	network := api.Network{
		ID: id,
	}

	err := s.db.First(&network).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, server.ErrRecordNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("could not retrieve network: %w", err)
	}

	return &network, nil
}

// Networks retrieves a list of all known networks.
func (s *Storage) Networks() ([]*api.Network, error) {

	var networks []*api.Network
	err := s.db.Find(&networks).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve networks: %w", err)
	}

	return networks, nil
}
