package api

import (
	"sync"

	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

type addressCache struct {
	*sync.Mutex
	addresses map[string][]identifier.Address
}

func newAddressCache() *addressCache {

	c := addressCache{
		Mutex:     &sync.Mutex{},
		addresses: make(map[string][]identifier.Address),
	}

	return &c
}

func (a *addressCache) get(id string) ([]identifier.Address, bool) {
	a.Lock()
	defer a.Unlock()

	address, ok := a.addresses[id]
	return address, ok
}

func (a *addressCache) set(id string, address []identifier.Address) {
	a.Lock()
	defer a.Unlock()

	a.addresses[id] = address
}
