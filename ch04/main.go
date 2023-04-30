package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}

func nonempty(strings []string) []string {
	var out []string
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
