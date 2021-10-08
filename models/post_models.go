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
func GetSinglePost(id string) views.Post {
	var post views.Post
	db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&post)
	return post
}

func UpdatePost(r *http.Request, id string) *mongo.UpdateResult {
	var post views.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		panic(err)
	}
	result, err := db.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{{"$set",
			bson.D{
				{"Title", post.Title},
				{"Content", post.Content},
			}},
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

func DeletePost(id string) *mongo.DeleteResult {
	result, err := db.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
	return result
}
