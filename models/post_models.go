package models

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"posts-crud-golang/database"
	"posts-crud-golang/views"
)

var db = database.GetCollection("posts")

func GetAllPosts() ([]views.Post, error) {
	var post views.Post
	var posts []views.Post
	cursor, err := db.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&post)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func CreatePost(r *http.Request) (*mongo.InsertOneResult, error) {
	var post views.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		return nil, err
	}
	post.ID = uuid.New().String()
	res, insertErr := db.InsertOne(context.TODO(), post)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	return res, nil
}
