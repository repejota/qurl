// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import "net/http"

// Request ...
type Request struct {
}

// NewRequest ...
func NewRequest() *Request {
	r := &Request{}
	return r
}

// Fetch ...
func (r *Request) Fetch(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
