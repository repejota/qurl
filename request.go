// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// HTTPClient ...
type HTTPClient interface {
	Fetch(url string) (*http.Response, error)
}

// Request represents the call being made to retrieve the contents of an URL.
type Request struct {
}

// NewRequest returns a new request instance.
func NewRequest() *Request {
	return &Request{}
}

// Fetch performs an HTTP GET call to anURL and fetch the contents.
func (r *Request) Fetch(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	return client.Do(req)
}

// FakeRequest ...
type FakeRequest struct {
}

// NewFakeRequest returns a new request instance.
func NewFakeRequest() *FakeRequest {
	return &FakeRequest{}
}

// Fetch performs an HTTP GET call to anURL and fetch the contents.
func (r *FakeRequest) Fetch(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	body := "Hello world"
	resp := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Request:       req,
		Header:        make(http.Header),
	}
	return resp, nil
}
