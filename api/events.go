package api

// FIXME: If it gets introduced, use the symbols from the indexer repo
// in order to not duplicate knowledge.

// EventType represents all possible NFT-related events.
type EventType int

const (
	Mint EventType = iota + 1
	Transfer
	Sale
	Burn
)

// String returns the string representation of a given event type
func (e EventType) String() string {

	switch e {
	case Mint:
		return "mint"

	case Transfer:
		return "transfer"

	case Sale:
		return "sale"

	case Burn:
		return "burn"

	default:
		return ""
	}
}
