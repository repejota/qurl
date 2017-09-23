// Copyright 2017 The qurl Authors. All rights reserved.

// Package routes implements all the HTTP entry points for this microservice.
package routes

import (
	"encoding/json"
	"net/http"

	"github.com/repejota/qurl"
)

// Query fetch an URL and returns JSON with the data obtained.
func Query(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	qurl := qurl.NewQURL()
	err := qurl.SetURL(queryParams.Get("url"))
	if err != nil {
		http.Error(w, "INVALID_URL", http.StatusBadRequest)
		return
	}

	response, err := qurl.Query(queryParams)
	if err != nil {
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
