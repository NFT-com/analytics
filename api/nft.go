package api

import (
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

// getNFT returns a single NFT based on its ID.
func (s *Server) getNFT(id string) (*api.NFT, error) {

	nft, err := s.storage.NFT(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	return nft, nil
}

// getNFTByTokenID returns a single NFT based on the combination of chainID, contract address and token ID.
func (s *Server) getNFTByTokenID(chainID string, contract string, tokenID string) (*api.NFT, error) {

	nft, err := s.storage.NFTByTokenID(chainID, contract, tokenID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	return nft, nil
}

// nfts returns a list of NFTs fitting the search criteria.
func (s *Server) nfts(owner *string, collection *string, rarityMin *float64, orderBy api.NFTOrder) ([]*api.NFT, error) {

	nfts, err := s.storage.NFTs(owner, collection, rarityMin, orderBy)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nfts: %w", err)
	}

	return nfts, nil
}
