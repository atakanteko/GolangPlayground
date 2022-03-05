package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies []Movie
	movie1 := Movie{
		Title:  "Saving Private Ryan",
		Year:   1998,
		Color:  true,
		Actors: []string{"Tom Hanks", "Matt Damon", "Tom Sizemore", "Vin Diesel"},
	}
	var actors []string
	actors = append(actors, "Henry Fonda", "Martin Balsam", "Joseph Sweeney")
	movie2 := Movie{
		Title:  "12 Angry Men",
		Year:   1957,
		Color:  false,
		Actors: actors,
	}
	movies = append(movies, movie1, movie2)
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON Marshalling failed: %s", err)
	}
	fmt.Printf("%s", data)

	var titles []struct{ Title string }
	err = json.Unmarshal([]byte(data), &titles)
	if err != nil {
		log.Println(err)
	}
	for _, title := range titles {
		fmt.Println(title.Title)
	}
}
