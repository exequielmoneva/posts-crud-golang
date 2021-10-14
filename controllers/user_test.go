package controllers_test

import (
	"net/http"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	_, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatalf("could not retrieve users: %v", err)
	}
}
