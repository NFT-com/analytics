package api

import (
	"github.com/NFT-com/graph-api/events/models/events"
)

// Storage retrieves the list of events given the selector and a token.
// Token is used to determine from where in the record set iteration should continue.
type Storage interface {
	Transfers(events.TransferSelector, string) ([]events.Transfer, string, error)
	Sales(events.SaleSelector, string) ([]events.Sale, string, error)
}
