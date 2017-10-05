package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParameters(t *testing.T) {
	testCases := []struct {
		params     map[string]string
		statusCode int
	}{
		{
			map[string]string{"delay": "0"},
			http.StatusOK,
		},
		{
			map[string]string{"delay": "2"},
			http.StatusOK,
		},
		{
			map[string]string{"delay": "1"},
			http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.params), func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/", nil)
			q := req.URL.Query()
			for k, v := range tc.params {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()
			rec := httptest.NewRecorder()
			root(rec, req)
			res := rec.Result()
			if res.StatusCode != tc.statusCode {
				t.Errorf("`%v` failed, got %v, expected %v", tc, res.StatusCode, tc.statusCode)
			}
		})
	}
}
