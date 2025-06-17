package graphs

func getTestGraph() *Graph[string] {
	graph := NewGraph[string]()
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "E")
	return graph
}
