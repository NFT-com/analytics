package storage

import (
	"errors"
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

// NFT returns a single NFT based on the ID.
func (s *Storage) NFT(id string) (*api.NFT, error) {

	nft := api.NFT{
		ID: id,
	}

	err := s.db.First(&nft).Error
	if err != nil {
		// FIXME: err not found is a separate thing
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	return &nft, nil
}

// NFTByTokenID returns a single NFT based on the chain, contract and the tokenID.
func (s *Storage) NFTByTokenID(chainID string, contract string, tokenID string) (*api.NFT, error) {

	if chainID == "" || contract == "" || tokenID == "" {
		return nil, errors.New("mandatory fields missing")
	}

	var nft api.NFT
	err := s.db.
		Joins("INNER JOIN collection c ON collection_id = c.id").
		Where("c.chain_id = ?", chainID).
		Where("c.address = ?", contract).
		Where("token_id = ?", tokenID).
		First(&nft).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	return &nft, nil
}

func (s *Storage) NFTs() ([]*api.NFT, error) {
	return nil, fmt.Errorf("TBD: not implemented")
}
