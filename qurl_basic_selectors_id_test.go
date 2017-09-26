// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIDSelectorNotPresent(t *testing.T) {
	targetURL := "https://www.example.com"
	requestURL := fmt.Sprintf("/q?url=%s&selector=.foo", targetURL)
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
	if len(response.Selectors[".foo"]) != 0 {
		t.Fatalf("Response selector '.foo' expected to have zero elements but got '%v'", len(response.Selectors[".foo"]))
	}
}

/*
func TestIDSelectorPresent(t *testing.T) {
	req, err := http.NewRequest("GET", "/q?url=http://localhost:6060&selector=%23idname", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(Query)

	handler.ServeHTTP(rec, req)

	var response Response
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Unmarshaling response failed %v", err)
	}

	if len(response.Selectors["#idname"]) != 1 {
		t.Fatalf("Response selector '#idname' expected to have one element but got '%v'", response.Selectors["#idname"])
	}

	if response.Selectors["#idname"][0].Text != "selector id content" {
		t.Fatalf("Response selector '#idname' expected to be 'selector id content' but got '%v'", response.Selectors["#idname"][0])
	}
}
*/
