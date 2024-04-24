package rss

import (
	"encoding/xml"
	"time"
)

type Item struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	PubDate     time.Time `xml:"pubDate"`
}

type Channel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	PubDate     time.Time `xml:"pubDate"`
	Items       []Item    `xml:"item"`
}

func GenerateFeed() ([]byte, error) {
	items := []Item{
		{
			Title:       "Article 1",
			Link:        "https://example.com/article1",
			Description: "This is the description of Article 1.",
			PubDate:     time.Now(),
		},
		{
			Title:       "Article 2",
			Link:        "https://example.com/article2",
			Description: "This is the description of Article 2.",
			PubDate:     time.Now().Add(-1 * time.Hour),
		},
	}

	feed := Channel{
		Title:       "Sample RSS Feed",
		Link:        "https://example.com",
		Description: "This is sample RSS feed generated using Golang.",
		PubDate:     time.Now(),
		Items:       items,
	}

	xmlData, err := xml.MarshalIndent(feed, "", "    ")
	if err != nil {
		return nil, err
	}

	rssFeed := []byte(xml.Header + string(xmlData))

	return rssFeed, nil
}
