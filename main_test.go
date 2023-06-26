package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	rec := httptest.NewRecorder()
	hello(rec, req)

	res := rec.Result()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read response: %v", err)
	}

	if !strings.Contains(string(b), "helloworld") {
		t.Fatal("\"helloworld\" missing")
	}
}
