package graphs

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		adjacency: make(map[T][]T),
	}
}

type Graph[T comparable] struct {
	adjacency map[T][]T
}

func (g *Graph[T]) AddNode(value T) *Graph[T] {
	if _, ok := g.adjacency[value]; !ok {
		g.adjacency[value] = make([]T, 0)
	}

	return g
}

func (g *Graph[T]) AddEdge(a, b T) *Graph[T] {
	g.AddNode(a).AddNode(b)

	g.adjacency[a] = append(g.adjacency[a], b)
	g.adjacency[b] = append(g.adjacency[b], a)
	return g
}

func (g *Graph[T]) GetNeighbors(value T) []T {
	if _, ok := g.adjacency[value]; !ok {
		return nil
	}

	return g.adjacency[value]
}

func (g *Graph[T]) Empty() bool {
	return len(g.adjacency) == 0
}
