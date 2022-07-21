package api

import (
	"github.com/NFT-com/analytics/graph/models/api"
)

type Storage interface {
	Network(id string) (*api.Network, error)
	Networks() ([]*api.Network, error)

	NFT(id string) (*api.NFT, error)
	NFTByTokenID(networkID string, contract string, tokenID string) (*api.NFT, error)
	NFTs(owner *string, collectionID *string, orderBy api.NFTOrder, limit uint, prefetchOwners bool) ([]*api.NFT, error)
	NFTTraits(id string) ([]api.Trait, error)
	NFTOwners(nftID string) ([]api.Owner, error)

	Collection(id string) (*api.Collection, error)
	CollectionByContract(networkID string, contract string) (*api.Collection, error)
	CollectionsByNetwork(networkID string) ([]*api.Collection, error)
	Collections(networkID *string, orderBy api.CollectionOrder) ([]*api.Collection, error)

	CollectionSize(id string) (uint, error)
	CollectionNFTs(collectionID string, limit uint, afterID string) (nfts []*api.NFT, more bool, err error)
	CollectionTraits(collectionID string) ([]api.Trait, error)
	CollectionOwners(collectionID string) (map[string][]api.Owner, error)

	MarketplaceCollections(marketplaceID string) ([]*api.Collection, error)
	MarketplaceNetworks(marketplaceID string) ([]*api.Network, error)
	MarketplacesForCollection(collectionID string) ([]*api.Marketplace, error)
	MarketplacesByNetwork(networkID string) ([]*api.Marketplace, error)
}
