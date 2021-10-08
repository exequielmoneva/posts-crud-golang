package controllers

import (
	"encoding/json"
	"net/http"
	"posts-crud-golang/models"
	"posts-crud-golang/views"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		posts, err := models.GetAllUsers()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(posts)
		return

	case "POST":
		post, err := models.CreateUser(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode("invalid request body")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("method not allowed")
		return
	}
}

func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	switch r.Method {
	case "GET":
		user := models.GetSingleUser(id)
		if user != (views.User{}) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("this author has not posted yet")
		return

	case "PUT":
		user := models.UpdateUser(r, id)
		if user != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("invalid request body")
		return

	case "DELETE":
		result := models.DeleteUser(id)
		if result.DeletedCount > 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("invalid id")
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("method not allowed")
		return
	}
}

func UserPostsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if r.Method == "GET" {
		posts, err := models.GetPostsFromUser(id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(posts)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("method not allowed")
		return
	}
}
