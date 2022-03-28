package api

// Chain represents the chain and its networks.
type Chain struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
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

// Marketplace represents a single NFT marketplace (e.g. Opensea, DefiKingdoms).
type Marketplace struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
}

// NFT represents a single Non-Fungible Token.
type NFT struct {
	ID           string  `json:"id"`
	TokenID      string  `json:"tokenID"`
	Owner        string  `json:"owner"`
	Rarity       float64 `json:"rarity"`
	CollectionID string  `json:"-"`
}
