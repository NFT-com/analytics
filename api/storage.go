package api

import (
	"github.com/NFT-com/events-api/models/events"
)

// FIXME: Change this once the interface becomes clear.
type Storage interface {
	Mints(MintSelector) ([]events.Mint, error)
	Transfers(TransferSelector) ([]events.Transfer, error)
	Sales(SaleSelector) ([]events.Sale, error)
	Burns(BurnSelector) ([]events.Burn, error)
}
