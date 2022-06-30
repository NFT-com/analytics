package query

import (
	"context"

	"github.com/NFT-com/analytics/graph/query"
)

// MarketplaceFields describes the paths to the relevant fields in a marketplace query.
type MarketplaceFields struct {
	Volume    string
	MarketCap string
	Sales     string
	Users     string
}

// Marketplace describes the marketplace query, specifically whether the GraphQL query requires
// fetching aggregated fields (stats) for the marketplace.
type Marketplace struct {
	Volume    bool
	MarketCap bool
	Sales     bool
	Users     bool
}

// ParseMarketplaceQuery parses the GraphQL query according to the specified configuration.
func ParseMarketplaceQuery(fields MarketplaceFields, ctx context.Context) *Marketplace {

	selection := query.GetSelection(ctx)

	query := Marketplace{
		Volume:    selection.Has(fields.Volume),
		MarketCap: selection.Has(fields.MarketCap),
		Sales:     selection.Has(fields.Sales),
		Users:     selection.Has(fields.Users),
	}

	return &query
}

// NeedStats returns true if any of the aggregated data is required.
func (q Marketplace) NeedStats() bool {
	return q.Volume || q.MarketCap || q.Sales || q.Users
}
