package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Post struct {
	Id          int    `json:"id"`
	Slug        string `json:"slug"`
	Url         string `json:"url"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Status      string `json:"status"`
	Category    string `json:"category"`
	PublishedAt string `json:"publishedAt"`
	UpdatedAt   string `json:"updatedAt"`
	UserId      int    `json:"userId"`
}

const URL = "https://jsonplaceholder.org/posts"

var httpClient = &http.Client{Timeout: 10 * time.Second}

// GetPosts returns a slice of Post structs
func GetPosts() ([]Post, error) {

	// Make the request to the API
	resp, err := httpClient.Get(URL)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}

	// Close the response body when the function returns
	defer resp.Body.Close()

	// Check if the status code is 200
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get posts: status code %d", resp.StatusCode)
	}

	// Read the body of the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body of posts response: %w", err)
	}

	// Convert the body into a slice of Post structs
	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert posts into struct: %w", err)
	}

	return posts, nil
}
