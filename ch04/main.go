// An insertion sort
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year  int    `json:"released"`
	Color bool   `json:color,omitempty"`
	Authors []string
}

var movies = []Movie{
	{Title: "Casablanca"},
	{Title: "Cool Hand Luke", Color: true},
	{Title: "Bullitt", Year: 1968, Color: true},
}

func main() {
	data, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatal("json.Marshal: %v", err)
	}
	fmt.Printf("%s\n", data)
}
