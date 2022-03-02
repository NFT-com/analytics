package storage

import (
	"github.com/NFT-com/indexer-api/models/api"
)

type Storage struct{}

func New() *Storage {

	storage := Storage{}
	return &storage
}

func (s *Storage) NFT(id string) (*api.Nft, error) {
	return &firstNft, nil
}

func (s *Storage) NFTs() ([]*api.Nft, error) {
	return sampleNfts, nil
}

func (s *Storage) Collection(id string) (*api.Collection, error) {
	return &firstCollection, nil
}

func (s *Storage) Collections() ([]*api.Collection, error) {
	return sampleCollections, nil
}
