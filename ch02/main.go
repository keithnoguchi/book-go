// A pop counter
package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Counts the number of 1s
func popCount(x uint64) int {
	return int(pc[byte(x>>(8*0))] +
		pc[byte(x>>(8*1))] +
		pc[byte(x>>(8*2))] +
		pc[byte(x>>(8*3))] +
		pc[byte(x>>(8*4))] +
		pc[byte(x>>(8*5))] +
		pc[byte(x>>(8*6))] +
		pc[byte(x>>(8*7))])
}

func main() {
	for n := uint64(0); n <= 100; n++ {
		fmt.Printf("%3d(0x%02x) has %2d 1s\n", n, n, popCount(n))
	}
}
