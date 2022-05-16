package api

// Lookup provides collection and marketplace address lookup based on the ID.
type Lookup interface {
	CollectionAddress(id string) (chainID uint, contractAddress string, err error)
}
