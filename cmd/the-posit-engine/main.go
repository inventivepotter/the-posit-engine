package main

import (
	"net/http"

	"the-posit-engine/cmd/the-posit-engine/app"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, app.Greetings)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
