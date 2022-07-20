package api

// These are some of the specific (expensive) fields for which we want to know
// whether they were requested before retrieving.
const (
	// NFT fields.
	fieldRarity       = "rarity"
	fieldTraits       = "traits"
	fieldOwners       = "owners"
	fieldPrice        = "trading_price"
	fieldAveragePrice = "average_price"

	// Collection NFT arguments.
	argumentFirst = "first"
	argumentAfter = "after"

	// Collection and marketplace fields.
	fieldNFTs      = "nfts"
	fieldVolume    = "volume"
	fieldMarketCap = "market_cap"
	fieldSales     = "sales"
	fieldUsers     = "users"
)
