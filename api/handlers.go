package api

import (
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

func (s *Server) NFT() (*api.Nft, error) {

	nft, err := s.storage.NFT("14")
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	return nft, nil
}

func (s *Server) Nfts() ([]*api.Nft, error) {

	nfts, err := s.storage.NFTs()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nfts: %w", err)
	}

	return nfts, nil
}

func (s *Server) Collection() (*api.Collection, error) {

	collection, err := s.storage.Collection("15")
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
