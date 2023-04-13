package feedstore

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/mmcdole/gofeed"
	"time"
)

type FeedStore struct {
	conn *pgx.Conn
}

func NewFeedStore(connString string) (*FeedStore, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return &FeedStore{conn: conn}, nil
}

func (fs *FeedStore) Close() {
	fs.conn.Close(context.Background())
}

func (fs *FeedStore) CreateTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS feed_items (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			author TEXT,
			published TIMESTAMP WITH TIME ZONE,
			link TEXT NOT NULL,
			categories TEXT[]
		)
	`
	_, err := fs.conn.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to create table: %v", err)
	}
	return nil
}

func (fs *FeedStore) SaveItem(item *gofeed.Item) error {
	query := `
		INSERT INTO feed_items (title, author, published, link, categories)
		VALUES ($1, $2, $3, $4, $5)
	`
	published, _ := time.Parse(time.RFC1123, item.Published)
	_, err := fs.conn.Exec(context.Background(), query, item.Title, item.Author.Name, published, item.Link, item.Categories)
	if err != nil {
		return fmt.Errorf("unable to save item: %v", err)
	}
	return nil
}
