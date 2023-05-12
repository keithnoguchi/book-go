// JSON encoding, p. 107
package main

import "fmt"

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title: "Cassablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	},
}

func main() {
	fmt.Println(movies)
}
