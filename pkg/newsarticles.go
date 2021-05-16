package pkg

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
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
