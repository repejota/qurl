// Copyright 2016-2017 The qurl Authors. All rights reserved.

// Package routes implements all the HTTP entry points for this microservice.
package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

// TeaPot is the dummy route responds with an HTTP 418 Teapot code.
func TeaPot(c echo.Context) error {
	return c.String(http.StatusTeapot, "I'm a teapot!")
}
