package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// TODO: Improve error handling - some errors need not be relayed to the user.
// See https://github.com/NFT-com/analytics/issues/12

// bindError is used when user input was malformed.
func bindError(err error) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

// apiError is used when something went wrong during request processing - e.g. the events couldn't
// be retrieved from the database.
func apiError(err error) *echo.HTTPError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
