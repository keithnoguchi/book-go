// Struct Embedding, p. 104
package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var c Circle
	var w Wheel
	fmt.Println(c, w)
}
