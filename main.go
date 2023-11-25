package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	defer timer("GetPosts")()
	posts, err := GetPosts()
	if err != nil {
		panic(err)
	}

	prettyStruct(posts[0])
}

// timer returns a function that prints the name argument and
// the elapsed time between the call to timer and the call to
// the returned function. The returned function is intended to
// be used in a defer statement:
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

// Pretty print a struct
func prettyStruct(intef interface{}) {
	output, _ := json.MarshalIndent(intef, "", "\t")
	fmt.Printf("%s \n", output)
}
