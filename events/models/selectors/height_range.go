package selectors

// HeightRange allows selecting events in a block range.
type HeightRange struct {
	StartHeight string `query:"start_height"`
	EndHeight   string `query:"end_height"`
}
