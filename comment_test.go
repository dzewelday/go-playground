package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetComment(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Write([]byte(`{"id": 1, "post_id": 1, "user_id": "1", "comment": "Test comment"}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use server URL as the API URL
	comment, err := GetComment(1)
	if err != nil {
		t.Errorf("Expected nil, got '%s'", err)
	}

	// Test the values
	if comment.Id != 1 {
		t.Errorf("Expected '1', got '%d'", comment.Id)
	}
	if comment.PostId != 1 {
		t.Errorf("Expected '1', got '%d'", comment.PostId)
	}
	if comment.UserId != "1" {
		t.Errorf("Expected '1', got '%s'", comment.UserId)
	}
	if comment.Comment != "Test comment" {
		t.Errorf("Expected 'Test comment', got '%s'", comment.Comment)
	}
}
