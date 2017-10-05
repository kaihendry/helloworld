package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParameters(t *testing.T) {
	testCases := map[string]struct {
		params     map[string]string
		statusCode int
	}{
		"fast": {
			map[string]string{"delay": "0"},
			http.StatusOK,
		},
		"slow": {
			map[string]string{"delay": "2"},
			http.StatusOK,
		},
	}

	for tc, tp := range testCases {
		req, _ := http.NewRequest("GET", "/", nil)
		q := req.URL.Query()
		for k, v := range tp.params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
		rec := httptest.NewRecorder()
		root(rec, req)
		res := rec.Result()
		if res.StatusCode != tp.statusCode {
			t.Errorf("`%v` failed, got %v, expected %v", tc, res.StatusCode, tp.statusCode)
		}
	}
}
