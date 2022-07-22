package api

// Network represents the blockchain networks.
// Mainnet and testnets of a specific blockchain are distinct network objects.
type Network struct {
	ID          string `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
}
