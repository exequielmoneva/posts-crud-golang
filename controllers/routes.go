package controllers

import (
	"encoding/json"
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Api is up")
	})

	// User routes
	mux.HandleFunc("/post", PostsHandler)
	mux.HandleFunc("/post/filter", PostIdHandler)
	return mux
}
