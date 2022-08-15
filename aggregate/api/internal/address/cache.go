package address

import (
	"sync"

	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// Cache provides concurrency-safe address cache. It provides a lookup of a blockchain address
// based on its ID.
type Cache struct {
	*sync.Mutex
	addresses map[string][]identifier.Address
}

// NewCache creates a new address Cache.
func NewCache() *Cache {

	c := &Cache{
		Mutex:     &sync.Mutex{},
		addresses: make(map[string][]identifier.Address),
	}

	return c
}

// Get retrieves the address from the cache.
func (a *Cache) Get(id string) ([]identifier.Address, bool) {
	a.Lock()
	defer a.Unlock()

	address, ok := a.addresses[id]
	return address, ok
}

// Set caches the address.
func (a *Cache) Set(id string, address []identifier.Address) {
	a.Lock()
	defer a.Unlock()

	a.addresses[id] = address
}
