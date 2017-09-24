// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// QURL is the main interface for the microservice.
type QURL struct {
}

// Query queries the URL and process all the data we want to fetch.
func (q *QURL) Query(rr HTTPClient, params url.Values) (*Response, error) {
	url := params.Get("url")
	response := NewResponse()
	response.URL = url
	// Fetch URL content
	resp, err := rr.Fetch(url)
	if err != nil {
		return response, err
	}
	response.Status = resp.StatusCode
	err = q.processHeaders(params, resp.Header, response)
	if err != nil {
		return nil, err
	}
	err = q.processSelectors(params, resp, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// processHeaders ...
func (q *QURL) processHeaders(params url.Values, headers http.Header, response *Response) error {
	for _, v := range params["header"] {
		response.Headers[v] = headers[v]
	}
	return nil
}

// processSelectors ...
func (q *QURL) processSelectors(params url.Values, resp *http.Response, response *Response) error {
	// Build a DOM from response
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return err
	}
	// Query the DOM with all selectors
	for _, v := range params["selector"] {
		// Process matching nodes
		doc.Find(v).Each(func(index int, selection *goquery.Selection) {
			// Node text
			element := &element{
				Text: selection.Text(),
			}
			// Node attributes
			for _, v := range selection.Nodes[0].Attr {
				attr := &attribute{
					Key:   v.Key,
					Value: v.Val,
				}
				element.Attributes = append(element.Attributes, attr)
			}
			response.Selectors[v] = append(response.Selectors[v], element)
		})
	}
	return nil
}
