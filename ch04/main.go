// encoding/json.MarshalIndent()
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title   string
	Year    int  `json:"released"`
	Color   bool `json:"color,omitempty"`
	Authors []string
}

func main() {
	movies := []Movie{
		{
			Title: "Casablanca", Year: 1942, Color: false,
			Authors: []string{"Humphrey Bogart", "Ingrid Bergman"},
		},
	}
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatal("marshal error: %s\n", err)
	}
	fmt.Printf("%s\n", data)
}
