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

func TestSetURL(t *testing.T) {
	expectedURL := "https://www.example.com"
	qurl := NewQURL()
	err := qurl.SetURL(expectedURL)
	if err != nil {
		t.Fatal(err)
	}
	if qurl.URL != expectedURL {
		t.Fatalf("QURL url expected to be %s but got %s", expectedURL, qurl.URL)
	}
	if qurl.Response.URL != expectedURL {
		t.Fatalf("QURL Response url expected to be %s but got %s", expectedURL, qurl.Response.URL)
	}
}

func TestSetURLInvalid(t *testing.T) {
	expectedURLInvalid := "foo"
	qurl := NewQURL()
	err := qurl.SetURL(expectedURLInvalid)
	if err.Error() != "parse foo: invalid URI for request" {
		t.Fatal(err)
	}
}
