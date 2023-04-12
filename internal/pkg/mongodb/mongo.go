package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartMongo() *mongo.Client {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASS") + "@" + os.Getenv("DATABASE_SERVER") + ":" + os.Getenv("DATABASE_PORT")))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	if err != nil {
		log.Fatal(err)
	}
	return mongoClient
}
