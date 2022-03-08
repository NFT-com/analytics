package api

import (
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

func (s *Server) NFT(id string) (*api.NFT, error) {

	nft, err := s.storage.NFT(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	return nft, nil
}

func (s *Server) Nfts() ([]*api.NFT, error) {

	nfts, err := s.storage.NFTs()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nfts: %w", err)
	}

	return nfts, nil
}

func (s *Server) Collection(id string) (*api.Collection, error) {

	collection, err := s.storage.Collection(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection: %w", err)
	}

	return collection, nil
}

func (s *Server) Collections() ([]*api.Collection, error) {

	collections, err := s.storage.Collections()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return collections, nil
}
