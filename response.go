// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

// Attribute represents a key, value pair of strrings fromn an HTML node
// property.
type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Element represents a simplified HTML element node. It only supports the node
// name and a list of attributes.
type Element struct {
	Text       string       `json:"text"`
	Attributes []*Attribute `json:"attributes"`
}

// Response represents the result struct received after querying an URL.
// Contains information about the URL and the data retrieved after proessing
// the content.
type Response struct {
	URL       string                `json:"url"`
	Status    int                   `json:"status"`
	Headers   map[string][]string   `json:"headers,omitempty"`
	Selectors map[string][]*Element `json:"selectors,omitempty"`
}

// NewResponse returns a new response instance.
func NewResponse() *Response {
	response := Response{
		Headers:   make(map[string][]string),
		Selectors: make(map[string][]*Element),
	}
	return &response
}
