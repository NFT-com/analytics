package api

import (
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")

	errRetrieveChainFailed       = errors.New("could not retrieve chain(s)")
	errRetrieveCollectionFailed  = errors.New("could not retrieve collection(s)")
	errRetrieveMarketplaceFailed = errors.New("could not retrieve marketplace(s)")
	errRetrieveNFTFailed         = errors.New("could not retrieve NFT(s)")
)
