package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://bazubu.com/cv-seminar-32943.html")
	if err != nil {
		log.Fatalln(err)
	}
	title := doc.Find("title").Text()
	fmt.Println(title)

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); name == "description" {
			description, _ := s.Attr("content")
			fmt.Printf("Description field: %s\n", description)
		}
	})
}
