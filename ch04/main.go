// A diercted graph
package main

import "fmt"

var graph map[string]map[string]float64

func main() {
	graph = make(map[string]map[string]float64)
	addEdge(graph, "A", "B")
	fmt.Println(hasEdge(graph, "A", "B"))
	fmt.Println(hasEdge(graph, "B", "A"))
}

func addEdge(graph map[string]map[string]float64, from, to string) {
	// assigning to the default map value will cause panic.
	if graph[from] == nil {
		graph[from] = make(map[string]float64)
	}
	graph[from][to] = 1.1
}

func hasEdge(graph map[string]map[string]float64, from, to string) float64 {
	// referencing the default map value won't cause panic.
	return graph[from][to]
}
