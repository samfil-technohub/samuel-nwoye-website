package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"regexp"

	"github.com/samfil-technohub/samuel-nwoye-website/controllers"
)

// instantiate the headless browser and function to test
func executeRequest(req *http.Request, function http.HandlerFunc) *httptest.ResponseRecorder {
	browser := httptest.NewRecorder()
	http.HandlerFunc(function).ServeHTTP(browser, req)

	return browser
}

// Check that the desired Status Code is Returned
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestIndex(t *testing.T) {
	// type of request, path and payload
	req, _ := http.NewRequest("GET", "/", nil)
	res := executeRequest(req, controllers.Index)

	checkResponseCode(t, http.StatusOK, res.Code)

	// instantiate regex on the returned response
	r, _ := regexp.Compile(res.Body.String())
 
	// Did our test pass
	actual := res.Body.String()
	expected := r.FindString("Welcome to my Homepage")
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}