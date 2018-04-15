package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	f, err := os.Open("./hoge.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".description").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}
