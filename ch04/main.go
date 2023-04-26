package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GRP
	RMB
	JPY
)

func main() {
	r := [...]int{1, 2, 3};
	fmt.Println(len(r))
	fmt.Println(r[2])
	fmt.Printf("%T\n", r)

	fmt.Println(JPY)
}
