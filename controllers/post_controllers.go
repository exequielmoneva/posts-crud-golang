package controllers

import (
	"encoding/json"
	"net/http"
	"posts-crud-golang/models"
	"posts-crud-golang/views"
)

func PostsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		posts, err := models.GetAllPosts()
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
		post, err := models.CreatePost(r)
		if err != nil {
			if err.Error() == "empty user_id" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode("author id must be specified")
				return
			}
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
func PostIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	switch r.Method {
	case "GET":
		post := models.GetSinglePost(id)
		if post != (views.Post{}) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(post)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("post not found")
		return

	case "PUT":
		post := models.UpdatePost(r, id)
		if post != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("invalid request body")
		return

	case "DELETE":
		result := models.DeletePost(id)
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
