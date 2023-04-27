package main

import "fmt"

// An in-place reverse function.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := make([]int, 20)
	for i := range a {
		a[i] = i
	}
	reverse(a)
	fmt.Println(a)
}
