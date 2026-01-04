package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	new_rrs := RSSFeed{}
	request, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &new_rrs, fmt.Errorf("Fail to create the new request: %w", err)
	}

	request.Header.Set("User-Agent", "gator")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return &new_rrs, fmt.Errorf("Fail to send the request: %w", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &new_rrs, fmt.Errorf("fail to read: %w", err)
	}
	err = xml.Unmarshal(body, &new_rrs)
	if err != nil {
		return &new_rrs, fmt.Errorf("fail to unMarsh: %w", err)
	}

	new_rrs.Channel.Title = html.UnescapeString(new_rrs.Channel.Title)
	new_rrs.Channel.Description = html.UnescapeString(new_rrs.Channel.Description)

	for i, item := range new_rrs.Channel.Item {
		new_rrs.Channel.Item[i].Title = html.UnescapeString(item.Title)
		new_rrs.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}

	return &new_rrs, nil
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
