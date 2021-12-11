package route

import (
	"movie-gateway/route/middleware"

	"github.com/labstack/echo/v4"
)

var middlewareHandler = map[string]echo.MiddlewareFunc{
	"app_key": middleware.AuthCheckAppKey,
}
