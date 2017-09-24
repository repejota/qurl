// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Fooo", "bar")
			fmt.Fprintf(w, `
				<html>
					<head>
						<title>Page Title</title>
					</head>
					<body>
						<div class="classname">selector class content</div>
						<div id="idname">selector id content</div>
					</body>
				</html>
			`)
		})
		log.Fatal(http.ListenAndServe(":6060", nil))
	}()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestNewRequest(t *testing.T) {
	request := NewRequest()
	if request == nil {
		t.Fatalf("Error on create an instance of Request")
	}
}

func TestFetch(t *testing.T) {
	request := NewRequest()
	status, _, _, err := request.Fetch("http://localhost:6060/")
	if err != nil {
		t.Fatal(err)
	}
	if status != 200 {
		t.Fatalf("Fetch url expected to return code '200' but got %d", status)
	}
}

func TestFetchFail(t *testing.T) {
	request := NewRequest()
	status, _, _, err := request.Fetch("http://invalidhost")
	if err.Error() != "Get http://invalidhost: dial tcp: lookup invalidhost: no such host" {
		t.Fatal(err)
	}
	if status != 500 {
		t.Fatalf("Fetch url expected to return code '500' but got %d", status)
	}
}
