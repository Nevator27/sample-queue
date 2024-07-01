package router

import (
	"net/http"

	rmq2 "github.com/Nevator27/simple-queue/rmq"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Broker *rmq2.Broker
	Server *echo.Echo
}

func (r *Router) Handler(ctx echo.Context) error {
	response := r.Broker.MakeRequest()
	if response.Error != nil {
		return ctx.String(http.StatusInternalServerError, response.Error.Error())
	}

	return ctx.String(http.StatusOK, response.Uuid)
}
