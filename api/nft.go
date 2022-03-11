package api

import (
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

// GetNFT returns a single NFT based on its ID.
func (s *Server) GetNFT(id string) (*api.NFT, error) {

	nft, err := s.storage.NFT(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	return nft, nil
}

// GetNFTByTokenID returns a single NFT based on the combination of chainID, contract address and token ID.
func (s *Server) GetNFTByTokenID(chainID string, contract string, tokenID string) (*api.NFT, error) {

	nft, err := s.storage.NFTByTokenID(chainID, contract, tokenID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	return nft, nil
}

// Nfts returns a list of NFTs fitting the search criteria.
func (s *Server) Nfts(owner *string, collection *string, rarityMin *float64, orderBy api.NFTOrder) ([]*api.NFT, error) {

	nfts, err := s.storage.NFTs(owner, collection, rarityMin, orderBy)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nfts: %w", err)
	}

	return nfts, nil
}
