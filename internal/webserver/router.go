package webserver

import "the-posit-engine/internal/handlers"

func (ws *webserver) bindRoutes() {
	ws.echo.GET("/", handlers.Notfound)
}
