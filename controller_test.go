package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFound(t *testing.T) {

	req, _ := http.NewRequest("GET", "/test/sam", nil)
	callDb = func(name string) MyData {
		return MyData{}
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getData)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestFound(t *testing.T) {

	req, _ := http.NewRequest("GET", "/test/sam", nil)
	callDb = func(name string) MyData {
		return MyData{Name: "sam"}
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getData)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
