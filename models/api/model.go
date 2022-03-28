package api

const (
	chainDBTable       = "chain"
	collectionDBTable  = "collection"
	nftDBTable         = "nft"
	marketplaceDBTable = "marketplace"
)

// Chain represents the chain and its networks.
type Chain struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// TableName returns the name of the underlying database table for the Chain.
func (c Chain) TableName() string {
	return chainDBTable
}

// Collection represents a group of NFTs that share the same smart contract.
type Collection struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Website     string `json:"website"`
	ImageURL    string `json:"image_url"`
	TokenURI    string `json:"token_uri"`
	ChainID     string `json:"-"`
}

// TableName returns the name of the underlying database table for the Collection.
func (c Collection) TableName() string {
	return collectionDBTable
}

// Marketplace represents a single NFT marketplace (e.g. Opensea, DefiKingdoms).
type Marketplace struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
}

// TableName returns the name of the underlying database table for the Marketplace.
func (m Marketplace) TableName() string {
	return marketplaceDBTable
}

// NFT represents a single Non-Fungible Token.
type NFT struct {
	ID           string  `json:"id"`
	TokenID      string  `json:"tokenID"`
	Owner        string  `json:"owner"`
	Rarity       float64 `json:"rarity"`
	CollectionID string  `json:"-"`
}

// TableName returns the name of the underlying database table for the NFT.
func (n NFT) TableName() string {
	return nftDBTable
}
