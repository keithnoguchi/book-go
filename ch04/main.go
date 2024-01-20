// Graph
package main

import "fmt"

type graph map[string]map[string]bool

func main() {
	g := newGraph()
	g.addEdge("a", "b")
	g.addEdge("b", "a")
	g.addEdge("a", "c")
	fmt.Println(g)
	fmt.Println(g.hasEdge("a", "b"))
	fmt.Println(g.hasEdge("a", "c"))
	fmt.Println(g.hasEdge("c", "a"))
}

func newGraph() graph {
	return make(map[string]map[string]bool)
}

func (g graph) addEdge(from, to string) {
	if g[from] == nil {
		g[from] = make(map[string]bool)
	}
	g[from][to] = true
}

func (g graph) hasEdge(from, to string) bool {
	return g[from][to]
}
