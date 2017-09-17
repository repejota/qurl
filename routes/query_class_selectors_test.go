// Copyright 2017 The qurl Authors. All rights reserved.

package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/repejota/qurl"
)

func TestClassSelectorNotPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "/q?url=http://localhost:6060&selector=.notpresent", nil)
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

	if len(response.Selectors[".notpresent"]) != 0 {
		t.Fatalf("Response selector '.notpresent' expected to have zero elements but got %v", len(response.Selectors[".notpresent"]))
	}
}

func TestClassSelectorPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "/q?url=http://localhost:6060&selector=.class", nil)
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

	if len(response.Selectors[".class"]) != 1 {
		t.Fatalf("Response selector '.class' expected to have one element but got %v", len(response.Selectors[".class"]))
	}
}
