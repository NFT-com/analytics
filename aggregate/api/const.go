package api

const (
	// idParam represents the name of the parameter used for collection/marketplace/NFT
	// ID in the API paths.
	idParam = "id"
)

// Available collection stats.
const (
	COLLECTION_VOLUME = iota + 1
	COLLECTION_MARKET_CAP
	COLLECTION_SALES
	COLLECTION_SIZE
	COLLECTION_AVERAGE_PRICE
	COLLECTION_FLOOR_PRICE
)
