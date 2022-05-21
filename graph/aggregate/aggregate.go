package aggregate

import (
	"net/url"

	"github.com/rs/zerolog"
)

// Client interacts with the aggregation API.
type Client struct {
	log    zerolog.Logger
	apiURL url.URL
}

// New creates a new aggregation API client.
func New(log zerolog.Logger, apiURL url.URL) *Client {

	c := Client{
		log:    log,
		apiURL: apiURL,
	}

	c.log.Debug().Str("api", apiURL.String()).Msg("creating aggregation API client")

	return &c
}
