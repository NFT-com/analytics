package api

import (
	"errors"

	"github.com/labstack/echo/v4"
)

// Burn returns all NFT burn events, according to the specified search criteria.
func (a *API) Burn(ctx echo.Context) error {

	var req Filter
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	events, err := a.storage.GetEvents(Burn, req)
	if err != nil {
		return apiError(err)
	}

	_ = events

	return errors.New("TBD: Not implemented")
}
