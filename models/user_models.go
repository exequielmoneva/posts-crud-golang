package models

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"posts-crud-golang/database"
	"posts-crud-golang/views"
)

var usersCollection = database.GetCollection("users")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GetAllUsers() ([]views.User, error) {
	var user views.User
	var users []views.User
	cursor, err := usersCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(r *http.Request) (*mongo.InsertOneResult, error) {
	var user views.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	user.ID = uuid.New().String()
	user.Password, _ = HashPassword(user.Password)
	res, insertErr := usersCollection.InsertOne(context.TODO(), user)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	return res, nil
}

func GetSingleUser(id string) views.User {
	var user views.User
	postsCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	return user
}

func GetPostsFromUser(id string) {

}

func UpdateUser() {

}

func DeleteUser() {

}
