package api

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
)

// Mint returns all NFT mint events, according to the specified search criteria.
func (a *API) Mint(ctx echo.Context) error {

	var req request
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	fmt.Printf("request: %+#v\n", req)

	return errors.New("TBD: Not implemented")
}
