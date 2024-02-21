// 5.6 Topological Sorting
package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"programming languages": {
		"data structures",
		"computer organization",
	},
	"operating systems": {
		"data structures",
		"algorithms",
	},
	"networks":         {"operating systems"},
	"formal languages": {"discrete math"},
	"discrete math":    {"intro to programming"},
	"data bases":       {"data structures"},
	"data structures":  {"discrete math"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"calculus": {"linear algebra"},
	"algorithms": {"data structures"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d\t%s\n", i+1, course)
	}
}

// Topological sorting, depth first example.
func topoSort(m map[string][]string) []string {
	seen := make(map[string]bool)
	var order []string
	var visit func([]string)
	visit = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visit(m[item])
				order = append(order, item)
			}
		}
	}
	// sort the keys of the prerequisite map to make the
	// result deterministic.
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visit(keys)
	return order
}
