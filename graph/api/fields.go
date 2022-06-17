package api

// These are some of the specific (expensive) fields for which we want to know
// whether they were requested before retrieving.
const (
	fieldRarity = "rarity"
	fieldTraits = "traits"
	fieldOwners = "owners"
	fieldNFTs   = "nfts"
)
