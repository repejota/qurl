// Copyright 2017 The qurl Authors. All rights reserved.

package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/repejota/qurl"
)

func TestHTTPHeaderNotPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "/q?url=http://localhost:6060&header=foobar", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(Query)

	handler.ServeHTTP(rec, req)

	var response qurl.Response
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Unmarshaling response failed %v", err)
	}

	if len(response.Headers["InvalidHeader"]) != 0 {
		t.Fatalf("Response header 'InvalidHeader' expected to be an empty slice but got %v", response.Headers["InvalidHeader"])
	}
}

func TestHTTPHeaderPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "/q?url=http://localhost:6060&header=Fooo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(Query)

	handler.ServeHTTP(rec, req)

	var response qurl.Response
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Unmarshaling response failed %v", err)
	}

	if len(response.Headers["Fooo"]) != 1 {
		t.Fatalf("Response header 'Fooo' expected to have one element but got %v", response.Headers["Fooo"])
	}
}
