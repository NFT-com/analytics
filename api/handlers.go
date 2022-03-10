package api

import (
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

func (s *Server) GetNFT(id string) (*api.NFT, error) {

	nft, err := s.storage.NFT(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	return nft, nil
}

func (s *Server) GetNFTByTokenID(chainID string, contract string, tokenID string) (*api.NFT, error) {

	nft, err := s.storage.NFTByTokenID(chainID, contract, tokenID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	return nft, nil
}

func (s *Server) Nfts(owner *string, collection *string, rarityMin *float64, orderBy api.NFTOrder) ([]*api.NFT, error) {

	nfts, err := s.storage.NFTs(owner, collection, rarityMin, orderBy)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nfts: %w", err)
	}

	return nfts, nil
}

func (s *Server) GetCollection(id string) (*api.Collection, error) {

	collection, err := s.storage.Collection(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection: %w", err)
	}

	return collection, nil
}

func (s *Server) GetCollectionByAddress(chainID string, contract string) (*api.Collection, error) {

	collection, err := s.storage.CollectionByAddress(chainID, contract)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection: %w", err)
	}

	return collection, nil
}

func (s *Server) Collections() ([]*api.Collection, error) {

	collections, err := s.storage.Collections()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return collections, nil
}

func (s *Server) GetCollectionNFTs(collectionID string) ([]*api.NFT, error) {

	nfts, err := s.storage.CollectionNFTs(collectionID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return nfts, nil
}

// GetChain retrieves a chain based on its ID.
func (s *Server) GetChain(id string) (*api.Chain, error) {

	chain, err := s.storage.Chain(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve chain: %w", err)
	}

	return chain, nil
}

// CollectionsByChain returns a list of collections on a given chain.
func (s *Server) CollectionsByChain(chainID string) ([]*api.Collection, error) {

	collections, err := s.storage.CollectionsByChain(chainID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return collections, nil
}

// CollectionListings returns a list of marketplaces where the collection is listed on.
func (s *Server) CollectionsListings(collectionID string) ([]*api.Marketplace, error) {

	marketplaces, err := s.storage.MarketplacesForCollection(collectionID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve marketplaces for a collection: %w", err)
	}

	return marketplaces, nil
}

// MarketplaceCollections returns a list of collections on a specified marketplace.
func (s *Server) MarketplaceCollections(marketplaceID string) ([]*api.Collection, error) {

	collections, err := s.storage.MarketplaceCollectionsList(marketplaceID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections on a marketplace: %w", err)
	}

	return collections, nil
}

// MarketplacesByChain returns a list of marketplaces on a specified chain.
func (s *Server) MarketplacesByChain(chainID string) ([]*api.Marketplace, error) {

	marketplaces, err := s.storage.MarketplacesByChain(chainID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve marketplaces: %w", err)
	}

	return marketplaces, nil
}

// MarketplaceChains returns a list of chains that collections listed on a marketplace reside on.
func (s *Server) MarketplaceChains(marketplaceID string) ([]*api.Chain, error) {

	chains, err := s.storage.MarketplaceChains(marketplaceID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve chains: %w", err)
	}

	return chains, nil
}
