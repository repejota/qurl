// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// QURL ...
type QURL struct {
	URL      string
	Request  *Request
	Response *Response
}

// NewQURL ...
func NewQURL() *QURL {
	qurl := &QURL{
		Request:  NewRequest(),
		Response: NewResponse(),
	}
	return qurl
}

// SetURL ...
func (q *QURL) SetURL(u string) error {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return err
	}
	q.URL = u
	q.Response.URL = q.URL
	return nil
}

// Query ...
func (q *QURL) Query(queryParams url.Values) (*Response, error) {
	// Fetch URL content
	r, err := q.Request.Fetch(q.URL)
	if err != nil {
		q.Response.Status = http.StatusInternalServerError
		return q.Response, err
	}
	defer func() { _ = r.Body.Close() }()

	err = q.processHeaders(queryParams, r)
	if err != nil {
		q.Response.Status = http.StatusInternalServerError
		return q.Response, err
	}

	err = q.processSelectors(queryParams, r)
	if err != nil {
		q.Response.Status = http.StatusInternalServerError
		return q.Response, err
	}

	q.Response.Status = r.StatusCode
	return q.Response, nil
}

// processHeaders ...
func (q *QURL) processHeaders(queryParams url.Values, response *http.Response) error {
	for _, v := range queryParams["header"] {
		q.Response.Headers[v] = response.Header[v]
	}
	return nil
}

// processSelectors ...
func (q *QURL) processSelectors(queryParams url.Values, response *http.Response) error {
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
