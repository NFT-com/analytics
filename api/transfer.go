package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/events-api/models/events"
)

// transferRequest describes a request to the transfer endpoint.
type transferRequest struct {
	events.TransferSelector
	Page string `query:"page"`
}

// transferResponse describes a response to the transfer listing request.
type transferResponse struct {
	Events   []events.Transfer `json:"events"`
	NextPage string            `json:"next_page,omitempty"`
}

// Transfer returns all NFT transfer events, according to the specified search criteria.
func (a *API) Transfer(ctx echo.Context) error {

	var req transferRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	transfers, token, err := a.storage.Transfers(req.TransferSelector, req.Page)
	if err != nil {
		return apiError(err)
	}

	res := transferResponse{
		Events:   transfers,
		NextPage: token,
	}

	return ctx.JSON(http.StatusOK, res)
}
