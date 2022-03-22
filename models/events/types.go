package events

// FIXME: If it gets introduced, use the symbols from the indexer repo
// in order to not duplicate knowledge.

// EventType represents all possible NFT-related events.
type EventType int

const (
	MintEvent EventType = iota + 1
	TransferEvent
	SaleEvent
	BurnEvent
)

// String returns the string representation of a given event type
func (e EventType) String() string {

	switch e {
	case MintEvent:
		return "mint"

	case TransferEvent:
		return "transfer"

	case SaleEvent:
		return "sale"

	case BurnEvent:
		return "burn"

	default:
		return ""
	}
}
