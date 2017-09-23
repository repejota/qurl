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
	queryParams := r.URL.Query()

	u := queryParams.Get("url")

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
	defer func() {
		err := response.Body.Close()
		if err != nil {
			http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
			return
		}
	}()

	err = processHeaders(queryParams, response, result)
	if err != nil {
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
	}

	err = processSelectors(queryParams, response, result)
	if err != nil {
		http.Error(w, "INTERNAL_ERROR", http.StatusInternalServerError)
		return
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

// processHeaders ...
func processHeaders(queryParams url.Values, response *http.Response, result *qurl.Response) error {
	for _, v := range queryParams["header"] {
		result.Headers[v] = response.Header[v]
	}
	return nil
}

// processSelectors ...
func processSelectors(queryParams url.Values, response *http.Response, result *qurl.Response) error {
	// Build a DOM from response content
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return err
	}

	// Query the DOM with all selectors
	for _, v := range queryParams["selector"] {

		// Process matching nodes
		doc.Find(v).Each(func(index int, selection *goquery.Selection) {

			// Node text
			element := &qurl.Element{
				Text: selection.Text(),
			}

			// Node attributes
			for _, v := range selection.Nodes[0].Attr {
				attr := &qurl.Attribute{
					Key:   v.Key,
					Value: v.Val,
				}
				element.Attributes = append(element.Attributes, attr)
			}

			result.Selectors[v] = append(result.Selectors[v], element)
		})
	}

	return nil
}
