package pkg

import (
	"io/ioutil"
	"net/http"
)

// Create the struct to convert XML to JSON

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
