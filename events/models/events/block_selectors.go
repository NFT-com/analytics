package events

// BlockSelector allows selecting events in a block range.
type BlockSelector struct {
	BlockStart string `query:"block_start"`
	BlockEnd   string `query:"block_end"`
}
