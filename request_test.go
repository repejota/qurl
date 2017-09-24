// Copyright 2017 The qurl Authors. All rights reserved.

package qurl

import (
	"testing"
)

func TestNewRequest(t *testing.T) {
	request := NewRequest()
	if request == nil {
		t.Fatalf("Error on create an instance of Request")
	}
}
