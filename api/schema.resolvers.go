package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/NFT-com/indexer-api/graph/generated"
	"github.com/NFT-com/indexer-api/models/api"
)

// FIXME: Fix this god-awful name - nFTServer

func (r *nFTServer) Collection(ctx context.Context, obj *api.NFT) (*api.Collection, error) {

	collection, err := r.Server.Collection(obj.CollectionID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection data: %w", err)
	}

	return collection, nil
}

func (r *queryServer) Nft(ctx context.Context, id string) (*api.NFT, error) {
	return r.Server.GetNFT(id)
}

func (r *queryServer) NftByTokenID(ctx context.Context, chainID string, contract string, tokenID string) (*api.NFT, error) {
	return nil, fmt.Errorf("TBD: not implemented")
}

func (r *queryServer) Nfts(ctx context.Context, owner *string, collection *string, rarityMin *float64, orderBy *api.NFTOrder) ([]*api.NFT, error) {
	return r.Server.Nfts()
}

func (r *queryServer) Collection(ctx context.Context, id string) (*api.Collection, error) {
	return r.Server.Collection(id)
}

func (r *queryServer) CollectionByAddress(ctx context.Context, chainID string, contract string) (*api.Collection, error) {
	return nil, fmt.Errorf("TBD: not implemented")
}

func (r *queryServer) Collections(ctx context.Context, chain *string, orderBy *api.CollectionOrder) ([]*api.Collection, error) {
	return r.Server.Collections()
}

// NFT returns generated.NFTResolver implementation.
func (r *Server) NFT() generated.NFTResolver { return &nFTServer{r} }

// Query returns generated.QueryResolver implementation.
func (r *Server) Query() generated.QueryResolver { return &queryServer{r} }

type nFTServer struct{ *Server }
type queryServer struct{ *Server }
