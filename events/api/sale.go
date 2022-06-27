package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/analytics/events/models/events"

	"github.com/NFT-com/analytics/events/models/selectors"
)

// saleRequest describes a request to the sales endpoint.
type saleRequest struct {
	selectors.SalesFilter
	Page string `query:"page"`
}

// saleResponse describes a response to the sale listing request.
type saleResponse struct {
	Sales    []events.Sale `json:"sales"`
	NextPage string        `json:"next_page,omitempty"`
}

// Sale returns all NFT sale events, according to the specified search criteria.
func (a *API) Sale(ctx echo.Context) error {

	var req saleRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	sales, token, err := a.storage.Sales(req.SalesFilter, req.Page)
	if err != nil {
		return apiError(err)
	}

	res := saleResponse{
		Sales:    sales,
		NextPage: token,
	}

	return ctx.JSON(http.StatusOK, res)
}
