package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/graph-api/events/models/events"
)

// mintRequest describes a request to the mints endpoint.
type mintRequest struct {
	events.MintSelector
	Page string `query:"page"`
}

// mintResponse describes a response to the mint listing request.
type mintResponse struct {
	Events   []events.Mint `json:"events"`
	NextPage string        `json:"next_page,omitempty"`
}

// Mint returns all NFT mint events, according to the specified search criteria.
func (a *API) Mint(ctx echo.Context) error {

	var req mintRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	mints, token, err := a.storage.Mints(req.MintSelector, req.Page)
	if err != nil {
		return apiError(err)
	}

	res := mintResponse{
		Events:   mints,
		NextPage: token,
	}

	return ctx.JSON(http.StatusOK, res)
}
