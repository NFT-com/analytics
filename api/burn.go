package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/events-api/models/events"
)

// burnRequest describes a request to the burns endpoint.
type burnRequest struct {
	events.BurnSelector
	Page string `query:"page"`
}

// burnResponse describes a response to the burn listing request.
type burnResponse struct {
	Events   []events.Burn `json:"events"`
	NextPage string        `json:"next_page,omitempty"`
}

// Burn returns all NFT burn events, according to the specified search criteria.
func (a *API) Burn(ctx echo.Context) error {

	var req burnRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	burns, token, err := a.storage.Burns(req.BurnSelector, req.Page)
	if err != nil {
		return apiError(err)
	}

	res := burnResponse{
		Events:   burns,
		NextPage: token,
	}

	return ctx.JSON(http.StatusOK, res)
}
