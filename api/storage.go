package api

import (
	"github.com/NFT-com/events-api/models/events"
)

// Storage retrieves the list of events given the selector and a token.
// Token is used to determine from where in the record set iteration should continue.
type Storage interface {
	Mints(events.MintSelector, string) ([]events.Mint, string, error)
	Transfers(events.TransferSelector, string) ([]events.Transfer, string, error)
	Sales(events.SaleSelector, string) ([]events.Sale, string, error)
	Burns(events.BurnSelector, string) ([]events.Burn, string, error)
}
