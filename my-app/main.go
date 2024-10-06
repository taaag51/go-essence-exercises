package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.bengo4.com/c_5/n_17912/"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 記事名の抽出
	articleTitle := doc.Find("h1.p-topics-article-header__title").Text()
	articleTitle = strings.TrimSpace(articleTitle)

	// 弁護士名の抽出
	lawyerName := doc.Find(".p-topics-supervisor__name").First().Text()
	lawyerName = strings.TrimSpace(lawyerName)

	fmt.Printf("記事名: %s\n", articleTitle)
	fmt.Printf("弁護士名: %s\n", lawyerName)
}
