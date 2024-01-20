// An insertion sort
package main

import (
	"fmt"
	"math/rand"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	// random numbers
	values := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		values = append(values, rand.Int()%100)
	}
	fmt.Printf("%v...\n", values[:16])
	sort(values)
	fmt.Printf("%v ... %v\n", values[0:8], values[len(values)-8:])
}

func sort(values []int) []int {
	var t *tree
	for _, v := range values {
		t = add(t, v)
	}
	appendValue(values[:0], t)
	return values
}

func appendValue(values []int, t *tree) []int {
	if t == nil {
		return values
	}
	values = appendValue(values, t.left)
	values = append(values, t.value)
	values = appendValue(values, t.right)
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
