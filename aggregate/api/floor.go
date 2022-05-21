package api

import (
	"errors"

	"github.com/labstack/echo/v4"
)

// FIXME: Think about this - the floor price for the entirety of the collection lifetime?
// Does that make sense?

// CollectionFloor handles the request for the floor price for NFTs in a collection.
func (a *API) CollectionFloor(ctx echo.Context) error {
	return errors.New("TBD: Not implemented")
}
