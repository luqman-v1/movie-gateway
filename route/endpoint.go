package route

import (
	"movie-gateway/domain/movie/handler"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Handle(c echo.Context) (err error)
}

var endpoint = map[string]Handler{
	"list_movie":   handler.NewList(),
	"detail_movie": handler.NewDetail(),
}
