package webserver

import (
	"github.com/labstack/echo/v4"
)

type webserver struct {
	echo *echo.Echo
}

var ws webserver

// Start is the webserver initiator
func Start() error {
	// initiate echo server
	ws.echo = echo.New()

	// add routes to the server
	ws.addRoutes()

	// start the server
	err := ws.echo.Start(":1323")
	return err
}
