package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Notfound serves 404
func Notfound(c echo.Context) error {
	return c.String(http.StatusNotFound, "Not Found")
}
