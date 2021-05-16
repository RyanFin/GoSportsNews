package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestViewAllMongoDBDatabases(t *testing.T) {
	// Instantiate a new client object
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://dbUser:GoSportsNews12@cluster0.1sizp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	// Instantiate a new context object
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	// connect to mongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	// List all databases currently available
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databases)
}

func TestDatabaseInsertion(t *testing.T) {
	// Instantiate a new client object
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://dbUser:GoSportsNews12@cluster0.1sizp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	// Instantiate a new context object
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	// connect to mongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	// Check the connection with a ping
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	/*
		------ Insert data into MongoDB -----
	*/
	// Access 'articles' collection within the mongoDB 'news' database
	collection := client.Database("news").Collection("articles")

	fmt.Println(collection)

	// insert record into mongoDB
	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.141596}})
	id := res.InsertedID
	fmt.Println("new record id: ", id)
}

func TestDeleteMultipleDocs(t *testing.T) {
	// Instantiate a new client object
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://dbUser:GoSportsNews12@cluster0.1sizp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	// Instantiate a new context object
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	// connect to mongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	// Check the connection with a ping
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	// Access 'articles' collection within the mongoDB 'news' database
	collection := client.Database("news").Collection("articles")

	// delete all where taxonomies eq ?
	collection.DeleteMany(ctx, bson.M{"taxonomies": "First Team"})
	collection.DeleteMany(ctx, bson.M{"taxonomies": "Club News"})
	collection.DeleteMany(ctx, bson.M{"taxonomies": "Video"})
	collection.DeleteMany(ctx, bson.M{"taxonomies": "Community"})
	collection.DeleteMany(ctx, bson.M{"taxonomies": "Players"})
	collection.DeleteMany(ctx, bson.M{"taxonomies": "Interviews"})
	collection.DeleteMany(ctx, bson.M{"taxonomies": "Brentford B Team"})
	collection.DeleteMany(ctx, bson.M{"taxonomies": "Galleries"})
	collection.DeleteMany(ctx, bson.M{"taxonomies": "Match Reports"})

}
