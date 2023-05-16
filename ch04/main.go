// encoding/json.MarshalIndent()
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title   string   `json:"title"`
	Year    int      `json:"released"`
	Color   bool     `json:"color,omitempty"`
	Authors []string `json:"authors"`
	comment string   // a private comment
}

func main() {
	movies := []Movie{
		{
			Title: "Casablanca", Year: 1942, Color: false,
			Authors: []string{"Humphrey Bogart", "Ingrid Bergman"},
		},
		{
			Title:   "Cool Hand Luke",
			Year:    1967,
			Authors: []string{"Paul Newman"},
			comment: "something internal",
		},
	}
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("marshal error: %s\n", err)
	}
	fmt.Printf("%s\n", data)
}
