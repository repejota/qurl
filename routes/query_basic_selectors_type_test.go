// Copyright 2017 The qurl Authors. All rights reserved.

package routes

/*
func TestBasicSelectorTypeNotPresent(t *testing.T) {
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

func TestBasicSelectorTypePresent(t *testing.T) {
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

	if response.Selectors["title"][0].Text != "Page Title" {
		t.Fatalf("Response selector 'title' expected to be 'Page Title' but got '%v'", response.Selectors["title"][0])
	}
}
*/
