package api

import (
	"errors"

	"github.com/labstack/echo/v4"
)

// CollectionVolume handles the request for the trading volume for a collection.
func (a *API) CollectionVolume(ctx echo.Context) error {
	return errors.New("TBD: Not implemented")
}

// MarketplaceVolume handles the request for the trading volume for a marketplace.
func (a *API) MarketplaceVolume(ctx echo.Context) error {
	return errors.New("TBD: Not implemented")
}
