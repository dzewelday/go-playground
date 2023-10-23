package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Comment struct {
	Id      int    `json:"id"`
	PostId  int    `json:"post_id"`
	UserId  string `json:"user_id"`
	Comment string `json:"comment"`
}

func GetComment(postId int) (Comment, error) {
	const url = "https://jsonplaceholder.org/comments/%d"

	res, err := http.Get(fmt.Sprintf(url, postId))
	if err != nil {
		return Comment{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return Comment{},
			fmt.Errorf("api call failed for comment %v with: %d", postId, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Comment{}, err
	}

	var comment Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		return Comment{}, err
	}

	return comment, nil
}
