package api

import (
	"errors"

	"github.com/labstack/echo/v4"
)

// Sale returns all NFT sale events, according to the specified search criteria.
func (a *API) Sale(ctx echo.Context) error {

	var req Filter
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	events, err := a.storage.GetEvents(Sale, req)
	if err != nil {
		return apiError(err)
	}

	_ = events

	return errors.New("TBD: Not implemented")
}
