package main

// Remeber how a graph can be represented using an adjacency matrix (which in JS we represented as an object in which each value was another object)
// In go, we can represent a graph using a map whose keys are strings and values are map[string]bool
// Meaning for every key that exists inside the inner map, it represents an edge with the key of outer map

var graph = make(map[string]map[string]bool)

func main() {

}

// Steps of adding an edge to a graph
// Using the from string, obtain the inner map
// If the inner map is nil, allocate a new map and assign it as a value to the graph
// then, in the inner map, add this new string and make it's value to be 'true'
func addEdge(from, to string) {
	edges := graph[from]

	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true

}

// Since the zero value of a bool type is always false, and one property of maps in Go is that if the element doesn't exist, they always return a zero value
// Therefore without checking anything, we can directly access the element and return the result
func hasEdge(from, to string) bool {
	return graph[from][to]
}
