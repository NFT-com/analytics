package currency

import (
	"sync"

	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// Cache maps and caches currency identifiers (chain ID and contract address pair) to our currency ID.
type Cache struct {
	*sync.Mutex
	ids map[identifier.Currency]string
}

// NewCache creates a new currency Cache.
func NewCache() *Cache {

	c := &Cache{
		Mutex: &sync.Mutex{},
		ids:   make(map[identifier.Currency]string),
	}

	return c
}

// Get retrieves the currency ID from the Cache.
func (c *Cache) Get(currency identifier.Currency) (string, bool) {
	c.Lock()
	defer c.Unlock()

	id, ok := c.ids[currency]
	return id, ok
}

// Set caches the currency ID.
func (c *Cache) Set(currency identifier.Currency, id string) {
	c.Lock()
	defer c.Unlock()

	c.ids[currency] = id
}
