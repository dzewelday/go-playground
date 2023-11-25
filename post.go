package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
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

const url string = "https://jsonplaceholder.org/posts"

func GetPosts() ([]Post, error) {

	// Make the request to the API
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("failed to get posts: " + err.Error())
	}

	// Close the response body when we're done
	defer resp.Body.Close()

	// Check if the status code is 200
	if resp.StatusCode != 200 {
		return nil, errors.New("api call for posts failed with status code: " + string(rune(resp.StatusCode)))
	}

	// Read the body of the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read body of posts response: " + err.Error())
	}

	// Convert the body into a slice of Post structs
	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, errors.New("failed to convert posts into struct: " + err.Error())
	}

	return posts, nil
}
