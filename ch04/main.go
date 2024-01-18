// 4.2.2. In-Place Slice Techniques
package main

import "fmt"

func main() {
	stack := []int{}
	stack = push(stack, 1)
	stack = push(stack, 99, 2020)
	fmt.Printf("%v\n", stack)
	for range stack {
		var x int
		stack, x = pop(stack)
		fmt.Printf("%v\n", x)
		fmt.Printf("%v\n", stack)
	}
	//stack, x := pop(stack) panic...
	//fmt.Printf("%v\n", x)
}

func push(stack []int, data ...int) []int {
	return append(stack, data...)
}

func pop(stack []int) ([]int, int) {
	x := stack[len(stack)-1]
	return stack[:len(stack)-1], x
}
