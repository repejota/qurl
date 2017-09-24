// Copyright 2017 The qurl Authors. All rights reserved.

package routes

import (
	"encoding/json"
	"net/http"

	"github.com/repejota/qurl"
)

// Query route fetch an URL and queries the response to get data.
func Query(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	qurl := qurl.NewQURL()

	// Query the target URL.
	response, err := qurl.Query(queryParams)
	if err != nil {
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	// Builds the response with the obtained data.
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	// Returns the response as JSON.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
