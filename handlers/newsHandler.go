package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewsHandler(rw http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	rw.WriteHeader(http.StatusOK)
	// fmt.Println("Hello news vars: %v", vars)

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

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var newsarticles []bson.M
	if err = cursor.All(ctx, &newsarticles); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(newsarticles)

	for i, e := range newsarticles {
		fmt.Println(i, ": ", e)
	}

	// Marshal data into a JSON string
	j, err := json.Marshal(newsarticles)

	// Write json as response output to client
	rw.Write(j)
}
