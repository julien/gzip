package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithHeaders(t *testing.T) {

	data, _ := ioutil.ReadFile("test.html")

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(data)
	})

	handler := GZip(next)
	r, _ := http.NewRequest("GET", "/", nil)
	r.Header.Set("Accept-Encoding", "deflate, gzip")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("got %v want 200", w.Code)
	}

	if w.Header().Get("Content-Type") == "" {
		t.Errorf("got %v want a value", w.Header().Get("Content-Type"))
	}

	if w.Header().Get("Content-Encoding") != "gzip" {
		t.Errorf("got %v want a gzip", w.Header().Get("Content-Encoding"))
	}
}

func TestWithoutHeaders(t *testing.T) {

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello"))
	})

	handler := GZip(next)
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("got %v want 200", w.Code)
	}

	if w.Header().Get("Content-Type") != "text/plain" {
		t.Errorf("got %v want a text/plain", w.Header().Get("Content-Type"))
	}

	if w.Header().Get("Content-Encoding") != "" {
		t.Errorf("got %v want nothing", w.Header().Get("Content-Encoding"))
	}
}
