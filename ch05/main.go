// 5.6 Anonymous Functions
package main

import "fmt"

func main() {
	fmt.Println("5.6 Anonymous Functions")
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
