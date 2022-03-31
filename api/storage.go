package api

import (
	"github.com/NFT-com/events-api/models/events"
)

// FIXME: Change this once the interface becomes clear.
type Storage interface {
	Mints(events.MintSelector, string) ([]events.Mint, string, error)
	Transfers(events.TransferSelector, string) ([]events.Transfer, string, error)
	Sales(events.SaleSelector, string) ([]events.Sale, string, error)
	Burns(events.BurnSelector, string) ([]events.Burn, string, error)
}
