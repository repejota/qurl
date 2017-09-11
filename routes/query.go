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
	URL     string              `json:"url"`
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers,omitempty"`
}

// NewResponse ...
func NewResponse() *Response {
	r := &Response{}
	r.Headers = make(map[string][]string)
	return r
}

// Query fetch an URL and returns JSON with the data obtained.
func Query(c echo.Context) error {
	queryParams := c.QueryParams()
	u := queryParams.Get("url")

	result := NewResponse()
	result.URL = u
	result.Status = http.StatusOK

	// Validate URL
	_, err := url.ParseRequestURI(u)
	if err != nil {
		result.Status = http.StatusBadRequest
		return c.JSON(result.Status, result)
	}

	// Fetch URL content
	response, err := http.Get(u)
	if err != nil {
		result.Status = http.StatusInternalServerError
		return c.JSON(result.Status, result)
	}

	// Process headers
	for _, v := range queryParams["header"] {
		result.Headers[v] = response.Header[v]
	}

	return c.JSON(result.Status, result)
}
