package api

import (
	"github.com/NFT-com/indexer/models/events"

	"github.com/NFT-com/analytics/events/models/selectors"
)

// Storage retrieves the list of events given the selector and a token.
// Token is used to determine from where in the record set iteration should continue.
type Storage interface {
	Transfers(selectors.TransferFilter, string) ([]events.Transfer, string, error)
	Sales(selectors.SalesFilter, string) ([]events.Sale, string, error)
}
