// Copyright 2017 The qurl Authors. All rights reserved.

// Package routes implements all the HTTP entry points for this microservice.

package routes

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo"
)

// Response is the type that defines a query result.
type Response struct {
	URL    string `json:"url"`
	Status int    `json:"status"`
}

// Query fetch an URL and returns JSON with the data obtained.
func Query(c echo.Context) error {
	result := &Response{}
	result.Status = http.StatusOK

	// Validate URL
	u := c.QueryParam("url")
	_, err := url.ParseRequestURI(u)
	if err != nil {
		result.Status = http.StatusBadRequest
	}

	result.URL = u
	return c.JSON(result.Status, result)
}
