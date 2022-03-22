package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// bindError is used when user input was malformed.
func bindError(err error) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

// apiError is used when something went wrong during request processing - e.g. the events couldn't
// be retrieved from the database.
func apiError(err error) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
