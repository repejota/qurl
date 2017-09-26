// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIDSelectorNotPresent(t *testing.T) {
	targetURL := "https://www.example.com"
	requestURL := fmt.Sprintf("/q?url=%s&selector=#foo", targetURL)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	q := &QURL{}
	freq := &FakeRequest{
		ExpectedBody:       `<p class="content">Page content</title>`,
		ExpectedStatusCode: http.StatusOK,
	}
	response, err := q.Query(freq, req.URL.Query())
	if err != nil {
		t.Fatal(err)
	}
	if response.Status != http.StatusOK {
		t.Errorf("response status expected to be %d but got %d", http.StatusOK, response.Status)
	}
	if response.URL != targetURL {
		t.Errorf("response url expected to be %s but got %s", targetURL, response.URL)
	}
	if len(response.Selectors["#foo"]) != 0 {
		t.Fatalf("Response selector '#foo' expected to have zero elements but got '%v'", len(response.Selectors["#foo"]))
	}
}

func TestIDSelectorPresent(t *testing.T) {
	targetURL := "https://www.example.com"
	requestURL := fmt.Sprintf("/q?url=%s&selector=#content", targetURL)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	q := &QURL{}
	freq := &FakeRequest{
		ExpectedBody:       `<p class="content">Page content</title>`,
		ExpectedStatusCode: http.StatusOK,
	}
	response, err := q.Query(freq, req.URL.Query())
	if err != nil {
		t.Fatal(err)
	}
	if response.Status != http.StatusOK {
		t.Errorf("response status expected to be %d but got %d", http.StatusOK, response.Status)
	}
	if response.URL != targetURL {
		t.Errorf("response url expected to be %s but got %s", targetURL, response.URL)
	}
	if len(response.Selectors["#content"]) != 1 {
		t.Fatalf("Response selector '#content' expected to have one element but got '%v'", len(response.Selectors["#content"]))
	}
	if response.Selectors["#content"][0].Text != "Page content" {
		t.Fatalf("Response selector '#content' expected to be %s but got %s", "Page content", response.Selectors["#content"][0].Text)
	}
}
