// 5.6 Topological Sorting
package main

import "fmt"

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal laguages",
		"computer organization",
	},
	"data structures": {"discrete math"},
	"data bases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal language": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {
		"data structures",
		"algorithms",
	},
	"programming languages": {
		"data structures",
		"computer organization",
	},
}

func main() {
	fmt.Printf("%#v\n", prereqs)
}
