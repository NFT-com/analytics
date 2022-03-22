package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func bindError(err error) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}
