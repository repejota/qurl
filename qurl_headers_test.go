// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHTTPHeaderNotPresent(t *testing.T) {
	targetURL := "https://www.example.com"
	requestURL := fmt.Sprintf("/q?url=%s&header=foobar", targetURL)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	q := &QURL{}
	freq := &FakeRequest{
		ExpectedBody:       "",
		ExpectedStatusCode: http.StatusOK,
	}
	response, err := q.Query(freq, req.URL.Query())
	if response.Status != http.StatusOK {
		t.Errorf("response status expected to be %d but got %d", http.StatusOK, response.Status)
	}
	if response.URL != targetURL {
		t.Errorf("response url expected to be %s but got %s", targetURL, response.URL)
	}
	if len(response.Headers["foobar"]) != 0 {
		t.Fatalf("Response header 'foobar' expected to be an empty slice but got %v", response.Headers["foobar"])
	}
}

func TestHTTPHeaderPresent(t *testing.T) {
	targetURL := "https://www.example.com"
	requestURL := fmt.Sprintf("/q?url=%s&header=Content-Type", targetURL)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	q := &QURL{}
	rheaders := make(http.Header)
	rheaders.Add("Content-Type", "text/html")
	freq := &FakeRequest{
		ExpectedBody:            "",
		ExpectedStatusCode:      http.StatusOK,
		ExpectedResponseHeaders: rheaders,
	}
	response, err := q.Query(freq, req.URL.Query())
	if response.Status != http.StatusOK {
		t.Errorf("response status expected to be %d but got %d", http.StatusOK, response.Status)
	}
	if response.URL != targetURL {
		t.Errorf("response url expected to be %s but got %s", targetURL, response.URL)
	}
	if len(response.Headers["Content-Type"]) != 1 {
		t.Fatalf("Response header 'Content-Type' expected to have one element got %v", len(response.Headers["Content-Type"]))
	}
	if response.Headers["Content-Type"][0] != "text/html" {
		t.Fatalf("Response header 'Content-Type' expected to be %s but got %s", "text/html", response.Headers["Content-Type"][0])
	}
}
