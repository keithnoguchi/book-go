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
	symbols := [...]string{USD: "$", EUR: "e", JPY: "y"}
	fmt.Println(symbols)
}
