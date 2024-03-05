// 5.10 Recover
package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f() (n int) {
	defer func() {
		n = recover().(int)
	}()
	panic(99)
}
