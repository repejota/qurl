// Copyright 2017 The qurl Authors. All rights reserved.

// Package routes implements all the HTTP entry points for this microservice.
package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestQuery(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query
	// parameters for now, so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/q?url=https://www.example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to
	// record the response.
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	err = Query(c)
	if err != nil {
		t.Fatalf("Query() failed %v", err)
	}

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "{\"url\":\"https://www.example.com\",\"status\":200}"
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), expected)
	}
}

func TestQueryInvalidURL(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query
	// parameters for now, so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/q?url=invalidurl", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to
	// record the response.
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	err = Query(c)
	if err != nil {
		t.Fatalf("Query() failed %v", err)
	}

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	expected := "{\"url\":\"invalidurl\",\"status\":400}"
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), expected)
	}
}
