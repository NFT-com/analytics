package api

// These are some of the specific (expensive) fields for which we want to know
// whether they were requested before retrieving.
const (
	rarityField = "rarity"
	traitField  = "traits"
	nftField    = "nfts"
)
