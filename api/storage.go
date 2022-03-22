package api

import (
	"github.com/NFT-com/events-api/models/events"
)

// FIXME: Change this once the interface becomes clear.
type Storage interface {
	Mints(filter Filter) ([]events.Mint, error)
	Transfers(filter Filter) ([]events.Transfer, error)
	Sales(filter Filter) ([]events.Sale, error)
	Burns(filter Filter) ([]events.Burn, error)
}
