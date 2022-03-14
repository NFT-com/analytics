package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/NFT-com/indexer-api/graph/generated"
	"github.com/NFT-com/indexer-api/models/api"
)

func (r *chainServer) Marketplaces(ctx context.Context, obj *api.Chain) ([]*api.Marketplace, error) {
	// Marketplaces handles expanding the list of Marketplaces within a Chain object.

	return r.Server.marketplacesByChain(obj.ID)
}

func (r *chainServer) Collections(ctx context.Context, obj *api.Chain) ([]*api.Collection, error) {
	// Collections handles expanding the list of Collections within a Chain object.

	return r.Server.collectionsByChain(obj.ID)
}

func (r *collectionServer) Chain(ctx context.Context, obj *api.Collection) (*api.Chain, error) {
	// Chain handles expanding the Chain object within a Collection object.

	return r.Server.getChain(obj.ChainID)
}

func (r *collectionServer) Marketplaces(ctx context.Context, obj *api.Collection) ([]*api.Marketplace, error) {
	// Marketplace handles expanding the list of Marketplaces within a Collection object.

	return r.Server.collectionsListings(obj.ID)
}

func (r *collectionServer) Nfts(ctx context.Context, obj *api.Collection) ([]*api.NFT, error) {
	// Nfts handles expanding the list of NFTs within a Collection object.

	return r.Server.getCollectionNFTs(obj.ID)
}

func (r *marketplaceServer) Chains(ctx context.Context, obj *api.Marketplace) ([]*api.Chain, error) {
	// Chains handles expanding the list of Chains within a Marketplace object.

	return r.Server.marketplaceChains(obj.ID)
}

func (r *marketplaceServer) Collections(ctx context.Context, obj *api.Marketplace) ([]*api.Collection, error) {
	// Collections handles expanding the list of Collections within a Marketplace object.

	return r.Server.marketplaceCollections(obj.ID)
}

func (r *nFTServer) Collection(ctx context.Context, obj *api.NFT) (*api.Collection, error) {
	// Collection handles expanding the Collection object within an NFT object.

	return r.Server.getCollection(obj.CollectionID)
}

func (r *queryServer) Chain(ctx context.Context, id string) (*api.Chain, error) {
	// Chain implements the `chain` GraphQL query

	return r.Server.getChain(id)
}

func (r *queryServer) Chains(ctx context.Context) ([]*api.Chain, error) {
	// Chains implements the `chains` GraphQL query.

	return r.Server.chains()
}

func (r *queryServer) Nft(ctx context.Context, id string) (*api.NFT, error) {
	// Nft implements the `nft` GraphQL query.

	return r.Server.getNFT(id)
}

func (r *queryServer) NftByTokenID(ctx context.Context, chainID string, contract string, tokenID string) (*api.NFT, error) {
	// NftByTokenID implements the `nftByTokenID` GraphQL query.

	return r.Server.getNFTByTokenID(chainID, contract, tokenID)
}

func (r *queryServer) Nfts(ctx context.Context, owner *string, collection *string, rarityMin *float64, orderBy *api.NFTOrder) ([]*api.NFT, error) {
	// Nfts implements the `nfts` GraphQL query.

	// FIXME: remove when all modes become supported
	switch orderBy.Field {

	case api.NFTOrderFieldValue:
		return nil, errors.New("TBD: sorting mode not supported")

	// supported sorting mode(s)
	case api.NFTOrderFieldCreationTime:
	case api.NFTOrderFieldRarity:
	}

	// NOTE: Ordering parameter is a pointer but gets initialized to the default value by the middleware.

	return r.Server.nfts(owner, collection, rarityMin, *orderBy)
}

func (r *queryServer) Collection(ctx context.Context, id string) (*api.Collection, error) {
	// Collection implements the `collection` GraphQL query.

	return r.Server.getCollection(id)
}

func (r *queryServer) CollectionByAddress(ctx context.Context, chainID string, contract string) (*api.Collection, error) {
	// CollectionByAddress implements the `collectionByAddress` GraphQL query.

	return r.Server.getCollectionByAddress(chainID, contract)
}

func (r *queryServer) Collections(ctx context.Context, chain *string, orderBy *api.CollectionOrder) ([]*api.Collection, error) {
	// Collections implements the `collections` GraphQL query.

	switch orderBy.Field {

	// FIXME: remove when all modes become supported
	default:
		return nil, errors.New("TBD: sorting mode not supported")

	// supported sorting mode(s)
	case api.CollectionOrderFieldCreationTime:
	}

	return r.Server.collections(chain, *orderBy)
}

// Chain returns generated.ChainResolver implementation.
func (r *Server) Chain() generated.ChainResolver { return &chainServer{r} }

// Collection returns generated.CollectionResolver implementation.
func (r *Server) Collection() generated.CollectionResolver { return &collectionServer{r} }

// Marketplace returns generated.MarketplaceResolver implementation.
func (r *Server) Marketplace() generated.MarketplaceResolver { return &marketplaceServer{r} }

// NFT returns generated.NFTResolver implementation.
func (r *Server) NFT() generated.NFTResolver { return &nFTServer{r} }

// Query returns generated.QueryResolver implementation.
func (r *Server) Query() generated.QueryResolver { return &queryServer{r} }

type chainServer struct{ *Server }
type collectionServer struct{ *Server }
type marketplaceServer struct{ *Server }
type nFTServer struct{ *Server }
type queryServer struct{ *Server }
