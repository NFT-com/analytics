package api

import (
	"errors"

	"github.com/labstack/echo/v4"
)

// CollectionCount handles the request for number of NFTs in a collection.
func (a *API) CollectionCount(ctx echo.Context) error {
	return errors.New("TBD: Not implemented")
}
