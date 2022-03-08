package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/NFT-com/indexer-api/graph/generated"
	"github.com/NFT-com/indexer-api/models/api"
)

func (r *queryServer) Nft(ctx context.Context, id string) (*api.Nft, error) {
	fields := getSelections(ctx)
	log.Printf("%+#v", fields)

	return r.Server.NFT(id)
}

func (r *queryServer) NftByTokenID(ctx context.Context, chainID string, contract string, tokenID string) (*api.Nft, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryServer) Nfts(ctx context.Context, owner *string, collection *string, rarityMin *float64, orderBy *api.NFTOrder) ([]*api.Nft, error) {
	return r.Server.Nfts()
}

func (r *queryServer) Collection(ctx context.Context, id string) (*api.Collection, error) {
	return r.Server.Collection(id)
}

func (r *queryServer) CollectionByAddress(ctx context.Context, chainID string, contract string) (*api.Collection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryServer) Collections(ctx context.Context, chain *string, orderBy *api.CollectionOrder) ([]*api.Collection, error) {
	return r.Server.Collections()
}

// Query returns generated.QueryResolver implementation.
func (r *Server) Query() generated.QueryResolver { return &queryServer{r} }

type queryServer struct{ *Server }
