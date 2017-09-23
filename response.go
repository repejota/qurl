// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

// Response is the type that defines a query result.
type Response struct {
	URL       string                `json:"url"`
	Status    int                   `json:"status"`
	Headers   map[string][]string   `json:"headers,omitempty"`
	Selectors map[string][]*Element `json:"selectors,omitempty"`
}

// Element represents an HTML element from a selector coincidence.
type Element struct {
	Text       string       `json:"text"`
	Attributes []*Attribute `json:"attributes"`
}

// Attribute ...
type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// NewResponse returns a response instance.
func NewResponse() *Response {
	r := &Response{}
	r.Headers = make(map[string][]string)
	r.Selectors = make(map[string][]*Element)
	return r
}
