// Copyright 2017 The qurl Authors. All rights reserved.

package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/repejota/qurl"
)

func TestHTTPHeaderNotPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "/q?url=https://www.example.com&header=foobar", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	err = Query(c)
	if err != nil {
		t.Fatalf("Query() failed %v", err)
	}

	var response qurl.Response
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Unmarshaling response failed %v", err)
	}

	if len(response.Headers["foobar"]) != 0 {
		t.Fatalf("Response header 'foobar' expected to be an empty slice but got %v", response.Headers["foobar"])
	}
}

func TestHTTPHeaderPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "/q?url=https://www.example.com&header=Content-Type", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	err = Query(c)
	if err != nil {
		t.Fatalf("Query() failed %v", err)
	}

	var response qurl.Response
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Unmarshaling response failed %v", err)
	}

	if len(response.Headers["Content-Type"]) != 1 {
		t.Fatalf("Response header 'Content-Type' expected to have one element but got %v", response.Headers["Content-Type"])
	}
}
