// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import (
	"io/ioutil"
	"net/http"
)

// Request ...
type Request struct {
}

// NewRequest ...
func NewRequest() *Request {
	r := Request{}
	return &r
}

// Fetch ...
func (r *Request) Fetch(url string) (int, *http.Header, []byte, error) {
	resp, err := http.Get(url)
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
