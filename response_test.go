// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import (
	"testing"
)

func TestNewResponse(t *testing.T) {
	result := NewResponse()
	if len(result.Headers) != 0 {
		t.Fatalf("Headers length expected to be 0 but got %d", len(result.Headers))
	}
	if len(result.Selectors) != 0 {
		t.Fatalf("Selectors length expected to be 0 but got %d", len(result.Selectors))
	}
}
