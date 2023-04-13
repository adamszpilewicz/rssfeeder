package main

import (
	"fmt"
	"rss/feedstore"
	"rss/rssfeeder"
)

func main() {
	// Example for Level Up Coding and printing the latest articles
	levelUpCodingReader := rssfeeder.NewRSSReader(
		"https://levelup.gitconnected.com", "")
	err := levelUpCodingReader.FetchFeed()
	if err != nil {
		fmt.Printf("Error fetching Level Up Coding RSS feed: %v\n", err)
	} else {
		fmt.Println("Level Up Coding:")
		levelUpCodingReader.PrintLatestArticles()
	}

	fmt.Println("--------------------------------------------------")

	// Example for a Medium publication and printing the latest articles
	mediumPublicationReader := rssfeeder.NewRSSReader(
		"https://medium.com", "better-programming")
	err = mediumPublicationReader.FetchFeed()
	if err != nil {
		fmt.Printf("Error fetching Medium publication RSS feed: %v\n", err)
	} else {
		fmt.Println("Better Programming:")
		mediumPublicationReader.PrintLatestArticles()
	}

	// Example for saving items to a database
	connString := "user=myuser password=mypassword host=localhost " +
		"port=5432 dbname=mydb sslmode=disable"
	fs, err := feedstore.NewFeedStore(connString)
	if err != nil {
		fmt.Printf("Error creating FeedStore: %v\n", err)
		return
	}
	defer fs.Close()

	err = fs.CreateTable()
	if err != nil {
		fmt.Printf("Error creating table: %v\n", err)
		return
	}

	for _, item := range levelUpCodingReader.Feed.Items {
		err = fs.SaveItem(item)
		if err != nil {
			fmt.Printf("Error saving item: %v\n", err)
		}
	}

	for _, item := range mediumPublicationReader.Feed.Items {
		err = fs.SaveItem(item)
		if err != nil {
			fmt.Printf("Error saving item: %v\n", err)
		}
	}

}
