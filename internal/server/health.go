package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func createHealth() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}
