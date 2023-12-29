package main

import "fmt"

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (b Book) String() string {
	return fmt.Sprintf("Title: %s\nAuthor: %s\nPages: %d\n", b.Title, b.Author, b.Pages)
}
