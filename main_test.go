package main

import (
	"net/http"
	"testing"
)

func TestIndexRoute(t *testing.T) {
	app := setupApp()

	req, _ := http.NewRequest("GET", "/", nil)
	resp, err := app.Test(req, -1)

	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}
}
