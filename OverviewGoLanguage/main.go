package main

import (
	"fmt"
	"github/atakanteko/golangplayground/github"
	"log"
)

func main() {
	var x int = 2
	var pointerToX *int
	pointerToX = &x
	fmt.Println(*pointerToX)
	*pointerToX = 5
	fmt.Println(*pointerToX)

	result, err := github.SearchIssues()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total Count: %d\n", result.TotalCount)
	fmt.Println(result.IncompleteResult)
}
