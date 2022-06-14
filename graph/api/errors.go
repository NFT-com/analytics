package api

import (
	"errors"

	"github.com/rs/zerolog"
)

var (
	ErrRecordNotFound = errors.New("record not found")

	errRetrieveNetworkFailed     = errors.New("could not retrieve network(s)")
	errRetrieveCollectionFailed  = errors.New("could not retrieve collection(s)")
	errRetrieveMarketplaceFailed = errors.New("could not retrieve marketplace(s)")
	errRetrieveNFTFailed         = errors.New("could not retrieve NFT(s)")
	errRetrieveTraitsFailed      = errors.New("could not retrieve NFT traits")
)

// TODO: Improve logging and error handling in the server code.
// See https://github.com/NFT-com/analytics/issues/7

// logError returns a prepared log entry. Errors resulting from requesting unknown records
// will be logged with a debug severity, in order to prevent excessive logging.
func (s *Server) logError(err error) *zerolog.Event {

	if errors.Is(err, ErrRecordNotFound) {
		return s.log.Debug().Err(err)
	}

	return s.log.Error().Err(err)
}
