package controllers_test

import (
	"net/http"
	"testing"
)

func TestGetAllPosts(t *testing.T) {
	_, err := http.NewRequest("GET", "/post", nil)
	if err != nil {
		t.Fatalf("could not retrieve posts: %v", err)
	}
}
