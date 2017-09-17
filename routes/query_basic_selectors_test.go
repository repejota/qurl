// Copyright 2017 The qurl Authors. All rights reserved.

package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/repejota/qurl"
)

func TestBasicSelectorNotPresent(t *testing.T) {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Fooo", "bar")
			fmt.Fprintf(w, `
				<html>
					<head>
						<title>Page Title</title>
					</head>
					<body>
						<div class="class">selector class content</div>
						<div id="id">selector id content</div>
					</body>
				</html>
			`)
		})
		http.ListenAndServe(":6060", nil)
	}()

	req, err := http.NewRequest("GET", "/q?url=http://localhost:6060&selector=title", nil)
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

	if len(response.Selectors["foo"]) != 0 {
		t.Fatalf("Response selector 'foo' expected to have zero elements but got '%v'", response.Selectors["foo"])
	}
}

func TestBasicSelectorPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "/q?url=http://localhost:6060&selector=title", nil)
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

	if len(response.Selectors["title"]) != 1 {
		t.Fatalf("Response selector 'title' expected to have one element but got '%v'", response.Selectors["title"])
	}

	if response.Selectors["title"][0] != "Page Title" {
		t.Fatalf("Response selector 'title' expected to be 'Page Title' but got '%v'", response.Selectors["title"][0])
	}
}
