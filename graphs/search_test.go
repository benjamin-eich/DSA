package graphs

import (
	"reflect"
	"testing"
)

func getSearchTestGraph() *Graph[string] {
	graph := NewGraph[string]()
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "E")
	return graph
}

func TestBfsReturnsNodesInBreadthFirstOrder(t *testing.T) {
	graph := getSearchTestGraph()
	expectedOrder := []string{"A", "B", "C", "D", "E"}

	order := Bfs(graph, "A")

	if !reflect.DeepEqual(order, expectedOrder) {
		t.Errorf("Bfs() = %v, want %v", order, expectedOrder)
	}
}

func TestBfsHandlesDisconnectedGraph(t *testing.T) {
	graph := NewGraph[string]()
	graph.AddEdge("A", "B")
	graph.AddEdge("C", "D")
	expectedOrder := []string{"A", "B"}

	order := Bfs(graph, "A")

	if !reflect.DeepEqual(order, expectedOrder) {
		t.Errorf("Bfs() = %v, want %v", order, expectedOrder)
	}

	expectedOrder = []string{"C", "D"}
	order = Bfs(graph, "C")

	if !reflect.DeepEqual(order, expectedOrder) {
		t.Errorf("Bfs() = %v, want %v", order, expectedOrder)
	}
}

func TestBfsReturnsSingleNodeForIsolatedStart(t *testing.T) {
	graph := NewGraph[string]()
	graph.AddEdge("A", "B")
	graph.AddEdge("C", "D")
	expectedOrder := []string{"E"}

	order := Bfs(graph, "E")

	if !reflect.DeepEqual(order, expectedOrder) {
		t.Errorf("Bfs() = %v, want %v", order, expectedOrder)
	}
}

func TestBfsHandlesEmptyGraph(t *testing.T) {
	graph := NewGraph[string]()
	expectedOrder := []string{}

	order := Bfs(graph, "A")

	if !reflect.DeepEqual(order, expectedOrder) {
		t.Errorf("Bfs() = %v, want %v", order, expectedOrder)
	}
}
