// package description
package main

var graph map[string]map[string]bool

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func main() {
	graph = make(map[string]map[string]bool)
}
