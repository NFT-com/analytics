package storage

// FIXME: temporary types for queries
type flatNFT struct {
	ID      string
	TokenID string
	Owner   string
	URI     string
	Rarity  float64
}

func (f flatNFT) TableName() string {
	return "nft"
}

type flatCollection struct {
	ID          string
	Name        string
	Description string
	Address     string
}

func (f flatCollection) TableName() string {
	return "collection"
}
