package api

import (
	"errors"

	"github.com/labstack/echo/v4"
)

// Transfer returns all NFT transfer events, according to the specified search criteria.
func (a *API) Transfer(ctx echo.Context) error {

	// FIXME: This code can be minimized since all endpoints are almost the same.
	// For now let's have it like this until we make sure there will be no event-specific filters.

	var req Filter
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	events, err := a.storage.GetEvents(Transfer, req)
	if err != nil {
		return apiError(err)
	}

	_ = events

	return errors.New("TBD: Not implemented")
}
