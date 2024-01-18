// 4.3. Maps
package main

import "fmt"

func main() {
	fmt.Println(
		equal(map[string]int{"keith": 0}, map[string]int{"mai": 0}),
	)
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
