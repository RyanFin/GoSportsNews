package pkg

import (
	"fmt"
	"testing"
)

func TestGetNewsArticles(t *testing.T) {
	resp, err := GetNewsArticles("50")
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	fmt.Println(resp)
}
