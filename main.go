package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	posts, err := GetPosts()
	if err != nil {
		log.Fatal(err)
	}

	err = prettyPrintStruct(posts[0])
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("GetPosts took %s", time.Since(start))
}

// prettyPrintStruct prints a struct in a pretty format
func prettyPrintStruct(strct any) error {
	output, err := json.MarshalIndent(strct, "", "\t")
	if err != nil {
		return fmt.Errorf("failed to marshal struct: %w", err)
	}

	fmt.Printf("%s \n", output)
	return nil
}
