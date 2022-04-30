package api

import (
	"github.com/NFT-com/graph-api/graph/models/api"
)

type Storage interface {
	Chain(id string) (*api.Chain, error)
	Chains() ([]*api.Chain, error)

	NFT(id string) (*api.NFT, error)
	NFTByTokenID(chainID string, contract string, tokenID string) (*api.NFT, error)
	NFTs(owner *string, collectionID *string, orderBy api.NFTOrder) ([]*api.NFT, error)
	NFTTraits(id string) ([]*api.Trait, error)

	Collection(id string) (*api.Collection, error)
	CollectionByContract(chainID string, contract string) (*api.Collection, error)
	CollectionNFTs(collectionID string) ([]*api.NFT, error)
	Collections(chain *string, orderBy api.CollectionOrder) ([]*api.Collection, error)
	CollectionsByChain(chainID string) ([]*api.Collection, error)
	CollectionTraits(collectionID string) ([]*api.Trait, error)
	CollectionSize(id string) (uint, error)

	MarketplaceCollections(marketplaceID string) ([]*api.Collection, error)
	MarketplaceChains(marketplaceID string) ([]*api.Chain, error)
	MarketplacesForCollection(collectionID string) ([]*api.Marketplace, error)
	MarketplacesByChain(chainID string) ([]*api.Marketplace, error)
}
