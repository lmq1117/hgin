package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net")

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("td").First().Text()
		//title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, band)
	})
}

func main() {
	ExampleScrape()
}
