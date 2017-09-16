// Copyright 2017 The qurl Authors. All rights reserved.
package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTeaPot(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query
	// parameters for now, so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/teapot", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to
	// record the response.
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(TeaPot)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusTeapot {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusTeapot)
	}

	// Check the response body is what we expect.
	expected := "I'm a Teapot!"
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), expected)
	}
}
