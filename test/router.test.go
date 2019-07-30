package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"shopstore/src/server/router"
	"testing"
)

//TestRouter -
func TestRouter(t *testing.T) {
	r := router.NewProductRouter()

	mockServer := httptest.NewServer(r)

	//define router
	res, err := http.Get(mockServer.URL + "/products")

	if err != nil {	
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status should be OK, got %v", res.StatusCode)
	}

	// In the next few lines, the response body is read, and converted to a string
	defer res.Body.Close()

	//read body
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	actual := string(body)
	expected := "Hello World!"

	if actual != expected {
		t.Errorf("Response should be %s, got %s", expected, actual)
	}
}
