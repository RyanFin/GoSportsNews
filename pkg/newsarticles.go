package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jasonlvhit/gocron"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Create the struct to convert XML to JSON
type NewListInformation struct {
	XMLName             xml.Name `xml:"NewListInformation"`
	Text                string   `xml:",chardata"`
	ClubName            string   `xml:"ClubName"`
	ClubWebsiteURL      string   `xml:"ClubWebsiteURL"`
	NewsletterNewsItems struct {
		Text               string `xml:",chardata"`
		NewsletterNewsItem []struct {
			Text              string `xml:",chardata"`
			ArticleURL        string `xml:"ArticleURL"`
			NewsArticleID     string `xml:"NewsArticleID"`
			PublishDate       string `xml:"PublishDate"`
			Taxonomies        string `xml:"Taxonomies"`
			TeaserText        string `xml:"TeaserText"`
			ThumbnailImageURL string `xml:"ThumbnailImageURL"`
			Title             string `xml:"Title"`
			OptaMatchId       string `xml:"OptaMatchId"`
		} `xml:"NewsletterNewsItem"`
	} `xml:"NewsletterNewsItems"`
}

func GetNewsArticles(count string) (string, error) {
	resp, err := http.Get(`https://www.brentfordfc.com/api/incrowd/getnewlistinformation?count=` + count + ``)
	if err != nil {
		return "", err
	}

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//Convert the body to type string
	sb := string(body)

	return sb, nil
}

func main() {
	// Run task every 5 seconds
	gocron.Every(2).Minutes().Do(newsletterFeed)
	// Start all the pending jobs
	<-gocron.Start()
}

func newsletterFeed() {
	/*
		------ Insert into Database ------
	*/

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

	var newListInfo NewListInformation

	// Get the response data for 50 news articles
	resp, err := GetNewsArticles("10")
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	err = xml.Unmarshal([]byte(resp), &newListInfo)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	// For each newsletter
	for _, newsLetter := range newListInfo.NewsletterNewsItems.NewsletterNewsItem {

		res, err := collection.InsertOne(ctx, newsLetter)
		if err != nil {
			fmt.Errorf("error: %v", err)
		}

		id := res.InsertedID

		fmt.Println("new record id: ", id)

	}
}
