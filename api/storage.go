package api

import (
	"github.com/NFT-com/events-api/models/events"
)

// FIXME: Change this once the interface becomes clear.
type Storage interface {
	Mints(events.MintSelector) ([]events.Mint, error)
	Transfers(events.TransferSelector) ([]events.Transfer, error)
	Sales(events.SaleSelector) ([]events.Sale, error)
	Burns(events.BurnSelector) ([]events.Burn, error)
}
