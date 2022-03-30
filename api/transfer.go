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

	// FIXME: This code can be minimized since all endpoints are almost the same.
	// For now let's have it like this until we make sure there will be no event-specific filters.

	var req transferRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	selector := req.TransferSelector
	if req.Page != "" {

		end, err := unpackPaginationToken(req.Page)
		if err != nil {
			return bindError(err)
		}

		selector.TimeSelector.End = end
	}

	transfers, err := a.storage.Transfers(selector)
	if err != nil {
		return apiError(err)
	}

	res := transferResponse{
		Events: transfers,
	}

	if len(transfers) > 0 {
		lastTimestamp := transfers[len(transfers)-1].Timestamp

		res.NextPage = createPaginationToken(lastTimestamp)
	}

	return ctx.JSON(http.StatusOK, res)
}
