package api

import (
	"github.com/NFT-com/indexer-api/models/api"
)

type Storage interface {
	Chain(id string) (*api.Chain, error)

	NFT(id string) (*api.NFT, error)
	NFTByTokenID(chainID string, contract string, tokenID string) (*api.NFT, error)
	NFTs(owner *string, collectionID *string, rarityMin *float64, orderBy api.NFTOrder) ([]*api.NFT, error)

	Collection(id string) (*api.Collection, error)
	CollectionByAddress(chainID string, contract string) (*api.Collection, error)
	Collections() ([]*api.Collection, error)
	CollectionsByChain(chainID string) ([]*api.Collection, error)
	CollectionNFTs(collectionID string) ([]*api.NFT, error)

	MarketplaceChains(marketplaceID string) ([]*api.Chain, error)
	MarketplacesByChain(chainID string) ([]*api.Marketplace, error)
	MarketplacesForCollection(collectionID string) ([]*api.Marketplace, error)
	MarketplaceCollectionsList(marketplaceID string) ([]*api.Collection, error)
}
