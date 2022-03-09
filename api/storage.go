package api

import (
	"github.com/NFT-com/indexer-api/models/api"
)

type Storage interface {
	NFT(id string) (*api.NFT, error)
	NFTs() ([]*api.NFT, error)
	Collection(id string) (*api.Collection, error)
	Collections() ([]*api.Collection, error)
}