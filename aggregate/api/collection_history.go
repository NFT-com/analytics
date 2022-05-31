package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionStatHistory returns the handler for the request stat.
func (a *API) CollectionStatHistory(mode int) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		// Unpack the collection request.
		request, err := a.unpackCollectionHistoryRequest(ctx)
		if err != nil {
			return err
		}

		// Retrieve the requested statistic.
		var value interface{}
		switch mode {

		case COLLECTION_VOLUME:
			value, err = a.stats.CollectionVolumeHistory(request.address, request.from, request.to)

		case COLLECTION_MARKET_CAP:
			value, err = a.stats.CollectionMarketCapHistory(request.address, request.from, request.to)

		case COLLECTION_SALES:
			value, err = a.stats.CollectionSalesHistory(request.address, request.from, request.to)

		case COLLECTION_FLOOR_PRICE:
			value, err = a.stats.CollectionFloorHistory(request.address, request.from, request.to)

		case COLLECTION_SIZE:
			value, err = a.stats.CollectionSizeHistory(request.address, request.from, request.to)

		case COLLECTION_AVERAGE_PRICE:
			value, err = a.stats.CollectionAverageHistory(request.address, request.from, request.to)

		// Invalid stat specified.
		default:
			err = fmt.Errorf("invalid stat (have: %v)", mode)
			return apiError(err)
		}

		if err != nil {
			err = fmt.Errorf("could not retrieve stat: %w", err)
			return apiError(err)
		}

		return ctx.JSON(http.StatusOK, value)
	}
}
