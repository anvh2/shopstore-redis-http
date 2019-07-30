package main

import (
	"net/http"
	"net/http/httptest"
	controller "shopstore/src/server/controller"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandleFunc(controller.GetAllProducts)

	hf.ServerHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("GetAllProducts returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `Hello world`
	actual := recorder.Body.String()

	if actual != expected {
		t.Errorf("GetAllProducts returned unexpected body: got %v want %v", actual, expected)
	}
}
