// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import (
	"io/ioutil"
	"net/http"
)

// Request represents the call being made to retrieve the contents of an URL.
type Request struct {
	URL string `json:"url"`
}

// NewRequest returns a new request instance.
func NewRequest() *Request {
	r := Request{}
	return &r
}

// Fetch performs an HTTP GET call to anURL and fetch the contents.
func (r *Request) Fetch() (int, *http.Header, []byte, error) {
	resp, err := http.Get(r.URL)
	if err != nil {
		return http.StatusInternalServerError, nil, nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return resp.StatusCode, nil, nil, err
	}
	return resp.StatusCode, &resp.Header, body, err
}
