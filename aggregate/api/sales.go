package api

import (
	"errors"

	"github.com/labstack/echo/v4"
)

// CollectionSales handles the request for number of sales for a collection.
func (a *API) CollectionSales(ctx echo.Context) error {
	return errors.New("TBD: Not implemented")
}

// MarketplaceSales handles the request for number of sales for a marketplace.
func (a *API) MarketplaceSales(ctx echo.Context) error {
	return errors.New("TBD: Not implemented")
}
