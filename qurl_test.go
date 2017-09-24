// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import "testing"

func TestNewQURL(t *testing.T) {
	expectedURL := ""
	qurl := NewQURL()
	if qurl.URL != "" {
		t.Fatalf("QURL url expected to be %s but got %s", expectedURL, qurl.URL)
	}
	if qurl.Response.URL != "" {
		t.Fatalf("QURL Response url expected to be %s but got %s", expectedURL, qurl.Response.URL)
	}
}
