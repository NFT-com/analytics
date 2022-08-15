package api

import (
	"sync"

	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// FIXME: Move the address caches to a separate package.

// currencyCache maps currency identifiers (chain ID and contract address pair) to our currency ID.
type currencyCache struct {
	*sync.Mutex
	ids map[identifier.Currency]string
}

func newCurrencyCache() *currencyCache {

	c := &currencyCache{
		Mutex: &sync.Mutex{},
		ids:   make(map[identifier.Currency]string),
	}

	return c
}

func (c *currencyCache) get(curr identifier.Currency) (string, bool) {
	c.Lock()
	defer c.Unlock()

	id, ok := c.ids[curr]
	return id, ok
}

func (c *currencyCache) set(curr identifier.Currency, id string) {
	c.Lock()
	defer c.Unlock()

	c.ids[curr] = id
}
