// Copyright 2017 The qurl Authors. All rights reserved.

package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/repejota/qurl"
)

func TestIDSelectorNotPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "/q?url=http://localhost:6060&selector=#unexistentid", nil)
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

	if len(response.Selectors["#unexistentid"]) != 0 {
		t.Fatalf("Response selector '#unexistentid' expected to have zero elements but got '%v'", response.Selectors["#unexistentid"])
	}
}

func TestIDSelectorPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "/q?url=http://localhost:6060&selector=#idname", nil)
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

	t.Log(response.Selectors)

	if len(response.Selectors["#idname"]) != 1 {
		t.Fatalf("Response selector '#idname' expected to have one element but got '%v'", response.Selectors["#idname"])
	}

	if response.Selectors["#idname"][0] != "selector id content" {
		t.Fatalf("Response selector '#idname' expected to be 'selector id content' but got '%v'", response.Selectors["#idname"][0])
	}
}
