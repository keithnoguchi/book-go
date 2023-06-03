// A web crawler
package main

import "fmt"

func main() {
	f := square
	f = negative
	fmt.Printf("%d\n", f(-9))
}

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}
