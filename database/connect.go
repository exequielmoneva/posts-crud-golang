package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func GetCollection(collection string) *mongo.Collection {
	mongoURI := "mongodb://" + os.Getenv("MONGO_HST") + ":" + os.Getenv("MONGO_PRT") + "/"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database("blog").Collection(collection)
}
