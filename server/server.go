// Copyright 2017 The qurl Authors. All rights reserved.

package server

import (
	"fmt"
	"net/http"

	"github.com/repejota/qurl/routes"
)

// QURLService ...
type QURLService struct {
	URL string
}

// Start starts the HTTP server for the qurl API microservice.
func Start(address string, port string) {

	http.HandleFunc("/teapot", routes.TeaPot)
	http.HandleFunc("/q", routes.Query)

	// Start server
	serveraddress := fmt.Sprintf("%s:%s", address, port)
	http.ListenAndServe(serveraddress, nil)
}
