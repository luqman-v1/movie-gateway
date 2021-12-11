package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/mwkosasih/soccer-gateway/util"
)

func AuthCheckAppKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		appKey := c.Request().Header.Get("app-key")
		if appKey != os.Getenv("app_key") {
			resp := util.Response{
				Code:    util.Unauthorized,
				Message: util.StatusMessage[util.InvalidAPIKey],
			}
			return resp.JSON(c)
		}
		return next(c)
	}
}
