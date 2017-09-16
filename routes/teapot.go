// Copyright 2017 The qurl Authors. All rights reserved.

// Package routes implements all the HTTP entry points for this microservice.
package routes

import (
	"io"
	"net/http"
)

// TeaPot is the dummy route responds with an HTTP 418 Teapot code.
func TeaPot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
	io.WriteString(w, "I'm a Teapot!")
}
