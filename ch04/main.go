package main

import "fmt"

func main() {
	r := [...]int{1, 2, 3};
	fmt.Println(len(r))
	fmt.Println(r[2])
	fmt.Printf("%T\n", r)
}
