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
	i := 0
	for _, s := range data {
		if s != "" {
			data[i] = s
			i++
		}
	}
	return data[:i]
}
