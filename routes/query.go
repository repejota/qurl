// Copyright 2017 The qurl Authors. All rights reserved.

// Package routes implements all the HTTP entry points for this microservice.
package routes

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// Query fetch an URL and returns JSON with the data obtained.
func Query(c echo.Context) error {
	url := c.QueryParam("url")

	log.Println(url)

	var result struct{}

	return c.JSON(http.StatusOK, result)
}
