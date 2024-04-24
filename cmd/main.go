package main

import (
	"fmt"
	"os"

	"github.com/matheus-alpe/go-rss-generator/internal/rss"
)

func main() {
	rssFeed, err := rss.GenerateFeed()
	if err != nil {
		fmt.Println("Error generating RSS feed:", err)
		return
	}

	file, err := os.Create("output/feed.xml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(rssFeed)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("RSS feed generated and saved to feed.xml")
}
