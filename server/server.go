// Copyright 2017 The qurl Authors. All rights reserved.

package server

import (
	"fmt"

	"bitbucket.org/ssgaas/api/routes"
	"github.com/labstack/echo"
)

// Start starts the HTTP server for the qurl API microservice.
func Start(address string, port string) {
	e := echo.New()
	e.DisableHTTP2 = true

	e.GET("/teapot", routes.TeaPot)

	// Start server
	serveraddress := fmt.Sprintf("%s:%s", address, port)
	e.Logger.Fatal(e.Start(serveraddress))
}
