package api

import (
	"github.com/NFT-com/graph-api/graph/models/api"
)

// getNFT returns a single NFT based on its ID.
func (s *Server) getNFT(id string) (*api.NFT, error) {

	nft, err := s.storage.NFT(id)
	if err != nil {
		s.logError(err).
			Str("id", id).
			Msg("could not retrieve nft")
		return nil, errRetrieveNFTFailed
	}

	return nft, nil
}

// getNFTByTokenID returns a single NFT based on the combination of chainID, contract address and token ID.
func (s *Server) getNFTByTokenID(chainID string, contract string, tokenID string) (*api.NFT, error) {

	nft, err := s.storage.NFTByTokenID(chainID, contract, tokenID)
	if err != nil {
		s.logError(err).
			Str("chain", chainID).
			Str("contract", contract).
			Str("token_id", tokenID).
			Msg("could not retrieve nft")
		return nil, errRetrieveNFTFailed
	}

	return nft, nil
}

// nfts returns a list of NFTs fitting the search criteria.
func (s *Server) nfts(owner *string, collection *string, rarityMin *float64, orderBy api.NFTOrder) ([]*api.NFT, error) {

	nfts, err := s.storage.NFTs(owner, collection, rarityMin, orderBy)
	if err != nil {
		log := s.logError(err)
		if owner != nil {
			log = log.Str("owner", *owner)
		}
		if collection != nil {
			log = log.Str("collection", *collection)
		}
		if rarityMin != nil {
			log = log.Float64("min_rarity", *rarityMin)
		}
		log.Msg("could not retrieve nfts")
		return nil, errRetrieveNFTFailed
	}

	return nfts, nil
}
