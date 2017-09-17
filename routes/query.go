// Copyright 2017 The qurl Authors. All rights reserved.

// Package routes implements all the HTTP entry points for this microservice.
package routes

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/repejota/qurl"
)

// Query fetch an URL and returns JSON with the data obtained.
func Query(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query().Get("url")

	result := qurl.NewResponse()
	result.URL = u
	result.Status = http.StatusOK

	// Validate URL
	_, err := url.ParseRequestURI(u)
	if err != nil {
		http.Error(w, "INVALID_URL", http.StatusBadRequest)
		return
	}

	// Fetch URL content
	response, err := http.Get(u)
	if err != nil {
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	queryParams := r.URL.Query()

	// Process headers
	for _, v := range queryParams["header"] {
		result.Headers[v] = response.Header[v]
	}

	// Process selectors
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}
	for _, v := range queryParams["selector"] {
		result.Selectors[v] = append(result.Selectors[v], doc.Find(v).Text())
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}
