// Copyright 2017 The qurl Authors. All rights reserved.
package qurl

import (
	"fmt"
	"net/http"
	"testing"
)

func TestQuery(t *testing.T) {
	targetURL := "https://www.example.com"
	requestURL := fmt.Sprintf("/q?url=%s", targetURL)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	q := &QURL{}
	freq := &FakeRequest{
		ExpectedBody:       "Hello world!",
		ExpectedStatusCode: http.StatusOK,
	}
	response, err := q.Query(freq, req.URL.Query())
	if response.Status != http.StatusOK {
		t.Errorf("response status expected to be %d but got %d", http.StatusOK, response.Status)
	}
	if response.URL != targetURL {
		t.Errorf("response url expected to be %s but got %s", targetURL, response.URL)
	}
}

func TestQueryInvalidURL(t *testing.T) {
	targetURL := "invalidurl"
	requestURL := fmt.Sprintf("/q?url=%s", targetURL)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	q := &QURL{}
	freq := &FakeRequest{
		ExpectedBody:       "",
		ExpectedStatusCode: http.StatusInternalServerError,
	}
	response, err := q.Query(freq, req.URL.Query())
	if response.Status != http.StatusInternalServerError {
		t.Errorf("response status expected to be %d but got %d", http.StatusInternalServerError, response.Status)
	}
	if response.URL != targetURL {
		t.Errorf("response url expected to be %s but got %s", targetURL, response.URL)
	}
}

func TestFailFetchURL(t *testing.T) {
	targetURL := "http://localhost"
	requestURL := fmt.Sprintf("/q?url=%s", targetURL)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	q := &QURL{}
	freq := &FakeRequest{
		ExpectedBody:       "",
		ExpectedStatusCode: http.StatusInternalServerError,
	}
	response, err := q.Query(freq, req.URL.Query())
	if response.Status != http.StatusInternalServerError {
		t.Errorf("response status expected to be %d but got %d", http.StatusInternalServerError, response.Status)
	}
	if response.URL != targetURL {
		t.Errorf("response url expected to be %s but got %s", targetURL, response.URL)
	}
}
