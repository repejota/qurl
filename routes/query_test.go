// Copyright 2017 The qurl Authors. All rights reserved.
package routes

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/repejota/qurl"
)

func TestQuery(t *testing.T) {
	targetURL := "http://localhost:6060"
	requestURL := fmt.Sprintf("/q?url=%s", targetURL)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	q := &qurl.QURL{
		URL: targetURL,
	}
	freq := &qurl.FakeRequest{}
	response, err := q.Query(freq, req.URL.Query())
	if response.Status != http.StatusOK {
		t.Errorf("response status expected to be %d but got %d", http.StatusOK, response.Status)
	}
	if response.URL != targetURL {
		t.Errorf("response url expected to be %s but got %s", targetURL, response.URL)
	}
}

/*
func TestQueryInvalidURL(t *testing.T) {
	req, err := http.NewRequest("GET", "/q?url=invalidurl", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(Query)

	handler.ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	// Check the response body is what we expect.
	expected := "INTERNAL_ERROR\n"
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), expected)
	}
}

func TestFailFetchURL(t *testing.T) {
	req, err := http.NewRequest("GET", "/q?url=http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(Query)

	handler.ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	// Check the response body is what we expect.
	expected := "INTERNAL_ERROR\n"
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), expected)
	}
}
*/
