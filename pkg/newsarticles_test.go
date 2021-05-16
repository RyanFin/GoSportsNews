package pkg

import (
	"encoding/xml"
	"fmt"
	"testing"
)

// func TestGetNewsArticles(t *testing.T) {
// 	resp, err := GetNewsArticles("50")
// 	if err != nil {
// 		fmt.Errorf("error: %v", err)
// 	}

// 	fmt.Println(resp)
// }

func TestUnmarshalXMLToStruct(t *testing.T) {
	var newListInfo NewListInformation

	resp, err := GetNewsArticles("50")
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	err = xml.Unmarshal([]byte(resp), &newListInfo)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	fmt.Println(newListInfo.ClubName)
	fmt.Println(newListInfo.ClubWebsiteURL)
	fmt.Println(newListInfo.NewsletterNewsItems)
	fmt.Println("-----")
	fmt.Println(newListInfo.NewsletterNewsItems.NewsletterNewsItem[0])
}
