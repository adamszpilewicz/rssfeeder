package rssfeeder

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

// RSSReader is a struct that contains the gofeed parser and the URL of the RSS Feed
// that we want to read.
type RSSReader struct {
	parser *gofeed.Parser
	url    string
	Feed   *gofeed.Feed
}

// NewRSSReader creates a new RSSReader struct.
func NewRSSReader(baseURL, slug string) *RSSReader {
	url := baseURL
	if slug != "" {
		url = baseURL + "/feed/" + slug
	} else {
		url = baseURL + "/feed"
	}
	return &RSSReader{
		parser: gofeed.NewParser(),
		url:    url,
	}
}

// FetchFeed fetches the RSS Feed and stores it in the RSSReader struct.
func (r *RSSReader) FetchFeed() error {
	feed, err := r.parser.ParseURL(r.url)
	if err != nil {
		return err
	}
	r.Feed = feed
	return nil
}

// PrintLatestArticles prints the latest articles from the RSS Feed.
func (r *RSSReader) PrintLatestArticles() {
	fmt.Printf("Latest articles:\n\n")
	for _, item := range r.Feed.Items {
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Author: %s\n", item.Author.Name)
		fmt.Printf("Published: %s\n", item.Published)
		fmt.Printf("Link: %s\n\n", item.Link)
		fmt.Printf("Categories: %v\n\n", item.Categories)
	}
}
