package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpoints(t *testing.T) {

	cases := []struct {
		method   string
		endpoint string
		handler  http.HandlerFunc
		wantCode int
	}{
		{"GET", "/", indexHandler(), http.StatusOK},
		{"GET", "/404", notFoundHandler(), http.StatusNotFound},
	}

	for _, c := range cases {
		r := httptest.NewRequest(c.method, c.endpoint, nil)
		w := httptest.NewRecorder()
		c.handler.ServeHTTP(w, r)
		want := c.wantCode
		got := w.Code
		if got != want {
			t.Errorf("want status = %d, got %d", want, got)
		}
	}
}
