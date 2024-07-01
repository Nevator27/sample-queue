package server

import (
	"log"

	rmq2 "github.com/Nevator27/simple-queue/rmq"

	"github.com/Nevator27/simple-queue/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(broker *rmq2.Broker) error {
	server := echo.New()
	server.Use(middleware.Recover())

	router := &router.Router{
		Server: server,
		Broker: broker,
	}

	server.GET("/", router.Handler)

	server.HideBanner = true
	server.HidePort = true

	log.Println("starting server...")

	return server.Start(":3001")
}
