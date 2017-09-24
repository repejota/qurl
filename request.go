// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import (
	"net/http"
)

// Request represents the call being made to retrieve the contents of an URL.
type Request struct {
}

// NewRequest returns a new request instance.
func NewRequest() *Request {
	r := Request{}
	return &r
}

// Fetch performs an HTTP GET call to anURL and fetch the contents.
func (r *Request) Fetch(url string) (*http.Response, error) {
	return http.Get(url)
}
