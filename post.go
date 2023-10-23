package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Post struct {
	Id       int    `json:"id"`
	Slug     string `json:"slug"`
	Url      string `json:"url"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Image    string `json:"image"`
	Status   int    `json:"status"`
	Category string `json:"category"`
	UserId   int    `json:"user_id"`

	PublishedAt int `json:"published_at"`
	CreatedAt   int `json:"created_at"`
}

const url string = "https://jsonplaceholder.org/posts"

func GetPosts() ([]Post, error) {

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("api call for posts was not successful")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("failed to read body of posts response")
	}

	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, errors.New("failed to convert posts into struct")
	}

	return posts, nil
}
