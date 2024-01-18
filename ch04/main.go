// 4.2.2. In-Place Slice Techniques
package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}

func nonempty(data []string) []string {
	out := data[:0]
	for _, s := range data {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
