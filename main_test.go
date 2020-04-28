package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"regexp"
)

func TestIndex(t *testing.T) {
	// type of request, path and payload
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// instantiate the headless browser
	rec := httptest.NewRecorder()

	// which function are we testing
	hf := http.HandlerFunc(Index)

	// make the request
	hf.ServeHTTP(rec, req)

	// Check the desired Status Code is Returned
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",status, http.StatusOK)
	}

	// instantiate regex on the returned response
	r, _ := regexp.Compile(rec.Body.String())
	
	// Did our test pass
	actual := rec.Body.String()
	expected := r.FindString("Welcome to my Homepage")
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}