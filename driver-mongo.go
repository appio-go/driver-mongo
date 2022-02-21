package driver_mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func Collection(db string, collection string) (*mongo.Collection, context.Context, *mongo.Client) {
	mongoURL, _ := os.LookupEnv("MONGO_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	getDatabase := client.Database(db)
	getCollection := getDatabase.Collection(collection)

	return getCollection, ctx, client
}
