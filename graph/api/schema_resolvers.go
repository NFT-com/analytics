package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/NFT-com/analytics/graph/generated"
	"github.com/NFT-com/analytics/graph/models/api"
)

func (r *collectionServer) Network(ctx context.Context, obj *api.Collection) (*api.Network, error) {
	// Network handles expanding the Network object within a Collection object.

	return r.Server.getNetwork(obj.NetworkID)
}

func (r *collectionServer) Marketplaces(ctx context.Context, obj *api.Collection) ([]*api.Marketplace, error) {
	// Marketplace handles expanding the list of Marketplaces within a Collection object.

	return r.Server.collectionsListings(ctx, obj.ID)
}

func (r *marketplaceServer) Networks(ctx context.Context, obj *api.Marketplace) ([]*api.Network, error) {
	// Networks handles expanding the list of Networks within a Marketplace object.

	return r.Server.marketplaceNetworks(obj.ID)
}

func (r *marketplaceServer) Collections(ctx context.Context, obj *api.Marketplace) ([]*api.Collection, error) {
	// Collections handles expanding the list of Collections within a Marketplace object.

	return r.Server.marketplaceCollections(ctx, obj.ID)
}

func (r *nFTServer) Collection(ctx context.Context, obj *api.NFT) (*api.Collection, error) {
	// Collection handles expanding the Collection object within an NFT object.

	return r.Server.getCollection(ctx, obj.Collection)
}

func (r *networkServer) Marketplaces(ctx context.Context, obj *api.Network) ([]*api.Marketplace, error) {
	// Marketplaces handles expanding the list of Marketplaces within a Network object.

	return r.Server.marketplacesByNetwork(ctx, obj.ID)
}

func (r *networkServer) Collections(ctx context.Context, obj *api.Network) ([]*api.Collection, error) {
	// Collections handles expanding the list of Collections within a Network object.

	return r.Server.collectionsByNetwork(ctx, obj.ID)
}

func (r *queryServer) Network(ctx context.Context, id string) (*api.Network, error) {
	// Network implements the `network` GraphQL query

	return r.Server.getNetwork(id)
}

func (r *queryServer) Networks(ctx context.Context) ([]*api.Network, error) {
	// Networks implements the `networks` GraphQL query.

	return r.Server.networks()
}

func (r *queryServer) Nft(ctx context.Context, id string) (*api.NFT, error) {
	// Nft implements the `nft` GraphQL query.

	return r.Server.getNFT(ctx, id)
}

func (r *queryServer) NftByTokenID(ctx context.Context, networkID string, contract string, tokenID string) (*api.NFT, error) {
	// NftByTokenID implements the `nftByTokenID` GraphQL query.

	return r.Server.getNFTByTokenID(ctx, networkID, contract, tokenID)
}

func (r *queryServer) Nfts(ctx context.Context, owner *string, collection *string, rarityMax *float64, orderBy *api.NFTOrder) ([]*api.NFT, error) {
	// Nfts implements the `nfts` GraphQL query.

	// TODO: Implement remaining sorting modes.
	// See https://github.com/NFT-com/analytics/issues/31
	switch orderBy.Field {

	case api.NFTOrderFieldValue, api.NFTOrderFieldRarity:
		return nil, errors.New("TBD: sorting mode not supported")

	// supported sorting mode(s)
	case api.NFTOrderFieldCreationTime:
	}

	// NOTE: Ordering parameter is a pointer but gets initialized to the default value by the middleware.

	return r.Server.nfts(ctx, owner, collection, rarityMax, *orderBy)
}

func (r *queryServer) Collection(ctx context.Context, id string) (*api.Collection, error) {
	// Collection implements the `collection` GraphQL query.

	return r.Server.getCollection(ctx, id)
}

func (r *queryServer) CollectionByAddress(ctx context.Context, networkID string, contract string) (*api.Collection, error) {
	// CollectionByAddress implements the `collectionByAddress` GraphQL query.

	return r.Server.getCollectionByContract(ctx, networkID, contract)
}

func (r *queryServer) Collections(ctx context.Context, networkID *string, orderBy *api.CollectionOrder) ([]*api.Collection, error) {
	// Collections implements the `collections` GraphQL query.

	// TODO: Implement remaining sorting modes.
	// See https://github.com/NFT-com/analytics/issues/31
	switch orderBy.Field {

	default:
		return nil, errors.New("TBD: sorting mode not supported")

	// supported sorting mode(s)
	case api.CollectionOrderFieldCreationTime:
	}

	return r.Server.collections(ctx, networkID, *orderBy)
}

// Collection returns generated.CollectionResolver implementation.
func (r *Server) Collection() generated.CollectionResolver { return &collectionServer{r} }

// Marketplace returns generated.MarketplaceResolver implementation.
func (r *Server) Marketplace() generated.MarketplaceResolver { return &marketplaceServer{r} }

// NFT returns generated.NFTResolver implementation.
func (r *Server) NFT() generated.NFTResolver { return &nFTServer{r} }

// Network returns generated.NetworkResolver implementation.
func (r *Server) Network() generated.NetworkResolver { return &networkServer{r} }

// Query returns generated.QueryResolver implementation.
func (r *Server) Query() generated.QueryResolver { return &queryServer{r} }

type collectionServer struct{ *Server }
type marketplaceServer struct{ *Server }
type nFTServer struct{ *Server }
type networkServer struct{ *Server }
type queryServer struct{ *Server }
