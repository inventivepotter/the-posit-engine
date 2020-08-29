package webserver

import "the-posit-engine/internal/handlers"

func (ws *webserver) addRoutes() {
	ws.echo.GET("/", handlers.Notfound)
}
