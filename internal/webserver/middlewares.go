package webserver

import "github.com/labstack/echo/v4/middleware"

var (
	defaultSecureConfig = middleware.SecureConfig{
		Skipper:            middleware.DefaultSkipper,
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "SAMEORIGIN",
	}
)

func (ws *webserver) bindMiddlewares() {
	ws.echo.Use(middleware.SecureWithConfig(defaultSecureConfig))
}
