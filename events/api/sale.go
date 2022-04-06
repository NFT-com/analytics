package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/graph-api/events/models/events"
)

// saleRequest describes a request to the sales endpoint.
type saleRequest struct {
	events.SaleSelector
	Page string `query:"page"`
}

// saleResponse describes a response to the sale listing request.
type saleResponse struct {
	Events   []events.Sale `json:"events"`
	NextPage string        `json:"next_page,omitempty"`
}

// Sale returns all NFT sale events, according to the specified search criteria.
func (a *API) Sale(ctx echo.Context) error {

	var req saleRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	sales, token, err := a.storage.Sales(req.SaleSelector, req.Page)
	if err != nil {
		return apiError(err)
	}

	res := saleResponse{
		Events:   sales,
		NextPage: token,
	}

	return ctx.JSON(http.StatusOK, res)
}
