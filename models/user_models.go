package models

import (
	"encoding/json"
	"net/http"
	"posts-crud-golang/database"
)

func Test (w http.ResponseWriter, r *http.Request){
	posts := database.GetCollection("posts")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts.Name())
	return
}
