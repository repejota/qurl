// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// QURL is the main ingterface for the microservice.
type QURL struct {
	URL      string
	Request  *Request
	Response *Response
}

// NewQURL creates an instance of QURL.
func NewQURL() *QURL {
	qurl := QURL{
		Request:  NewRequest(),
		Response: NewResponse(),
	}
	return &qurl
}

// SetURL validates and sets the target URL to be queried.
func (q *QURL) SetURL(u string) error {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return err
	}
	q.URL = u
	q.Request.URL = q.URL
	q.Response.URL = q.URL
	return nil
}

// Query queries the URL and process all the data we want to fetch.
func (q *QURL) Query(queryParams url.Values) error {
	// Fetch URL content
	statuscode, headers, body, err := q.Request.Fetch()
	if err != nil {
		q.Response.Status = statuscode
		return err
	}

	err = q.processHeaders(queryParams, *headers)
	if err != nil {
		q.Response.Status = statuscode
		return err
	}

	err = q.processSelectors(queryParams, bytes.NewReader(body))
	if err != nil {
		q.Response.Status = statuscode
		return err
	}

	q.Response.Status = statuscode
	return nil
}

// processHeaders ...
func (q *QURL) processHeaders(queryParams url.Values, headers http.Header) error {
	for _, v := range queryParams["header"] {
		q.Response.Headers[v] = headers[v]
	}
	return nil
}

// processSelectors ...
func (q *QURL) processSelectors(queryParams url.Values, body io.Reader) error {
	// Build a DOM from response content
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return err
	}

	// Query the DOM with all selectors
	for _, v := range queryParams["selector"] {

		// Process matching nodes
		doc.Find(v).Each(func(index int, selection *goquery.Selection) {

			// Node text
			element := &Element{
				Text: selection.Text(),
			}

			// Node attributes
			for _, v := range selection.Nodes[0].Attr {
				attr := &Attribute{
					Key:   v.Key,
					Value: v.Val,
				}
				element.Attributes = append(element.Attributes, attr)
			}

			q.Response.Selectors[v] = append(q.Response.Selectors[v], element)
		})
	}

	return nil
}
