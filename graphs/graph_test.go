package graphs

import (
	"reflect"
	"testing"
)

func TestAddNodeAddsNewNodeToGraph(t *testing.T) {
	graph := NewGraph[string]()
	graph.AddNode("A")

	if len(graph.adjacency) != 1 || len(graph.adjacency["A"]) != 0 {
		t.Errorf("AddNode() did not add the node correctly")
	}
}

func TestAddNodeDoesNotDuplicateExistingNode(t *testing.T) {
	graph := NewGraph[string]()
	graph.AddNode("A").AddNode("A")

	if len(graph.adjacency) != 1 {
		t.Errorf("AddNode() duplicated an existing node")
	}
}

func TestAddEdgeAddsEdgeBetweenNodes(t *testing.T) {
	graph := NewGraph[string]()
	graph.AddEdge("A", "B")

	if len(graph.adjacency["A"]) != 1 || graph.adjacency["A"][0] != "B" {
		t.Errorf("AddEdge() did not add the edge correctly for node A")
	}
	if len(graph.adjacency["B"]) != 1 || graph.adjacency["B"][0] != "A" {
		t.Errorf("AddEdge() did not add the edge correctly for node B")
	}
}

func TestAddEdgeCreatesNodesIfNotExist(t *testing.T) {
	graph := NewGraph[string]()
	graph.AddEdge("A", "B")

	if len(graph.adjacency) != 2 {
		t.Errorf("AddEdge() did not create missing nodes")
	}
}

func TestGetNeighborsReturnsCorrectNeighbors(t *testing.T) {
	graph := NewGraph[string]()
	graph.AddEdge("A", "B").AddEdge("A", "C")

	neighbors := graph.GetNeighbors("A")
	expected := []string{"B", "C"}

	if !reflect.DeepEqual(neighbors, expected) {
		t.Errorf("GetNeighbors() = %v, want %v", neighbors, expected)
	}
}

func TestGetNeighborsReturnsNilForNonExistentNode(t *testing.T) {
	graph := NewGraph[string]()

	neighbors := graph.GetNeighbors("A")

	if neighbors != nil {
		t.Errorf("GetNeighbors() = %v, want nil", neighbors)
	}
}

func TestEmptyReturnsTrueForEmptyGraph(t *testing.T) {
	graph := NewGraph[string]()

	if !graph.Empty() {
		t.Errorf("Empty() = false, want true")
	}
}

func TestEmptyReturnsFalseForNonEmptyGraph(t *testing.T) {
	graph := NewGraph[string]()
	graph.AddNode("A")

	if graph.Empty() {
		t.Errorf("Empty() = true, want false")
	}
}
