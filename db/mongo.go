package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var FileCollection *mongo.Collection

func Connect() {

	// url := ""
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		// context.TODO(),
		ctx,
		options.Client().ApplyURI("mongodb://localhost:27017/upload"),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	log.Println("DB connectred")
	FileCollection = client.Database("upload").Collection("meta")
}
