package main

import "fmt"

func main() {
	var x []int

	x = appendInt(x, []int{1, 2, 3}...)
	x = appendInt(x, 4)
	x = appendInt(x, 5, 6, 7)
	x = appendInt(x, x...)
	fmt.Println(x)
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x) // at least twice as big than x
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	fmt.Println(len(z), cap(z))
	fmt.Println(y)
	copy(z[len(x):], y)
	return z
}
