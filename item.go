package main

import (
	"github.com/PuerkitoBio/goquery"
)

type Item struct {
	Url   string
	Title string
	Body  string
}

func NewItem(url string) (*Item, error) {
	item := &Item{url, "", ""}

	doc, err := goquery.NewDocument(item.Url)
	if err != nil {
		return nil, err
	}
	item.Title = doc.Find("title").Text()
	item.Body = doc.Find("body").Text()
	return item, nil
}
