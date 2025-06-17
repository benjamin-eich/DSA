package graphs

import "github.com/benjamin-eich/DSA/queues"

func Bfs[T comparable](graph *Graph[T], start T) []T {
	visited := make(map[T]bool)
	queue := queues.NewQueue[T]()
	queue.Push(start)
	order := make([]T, 0)

	if graph.Empty() {
		return order
	}

	for !queue.Empty() {
		node, _ := queue.PopLeft()
		if _, ok := visited[node]; ok {
			continue
		}

		visited[node] = true
		order = append(order, node)

		for _, neighbor := range graph.GetNeighbors(node) {
			if _, ok := visited[neighbor]; !ok {
				queue.Push(neighbor)
			}
		}
	}

	return order
}

func Dfs[T comparable](graph *Graph[T], start T, visited map[T]bool) []T {
	if visited == nil {
		visited = make(map[T]bool)
	}

	order := make([]T, 0)

	if graph.Empty() {
		return order
	}

	if _, ok := visited[start]; !ok {
		visited[start] = true
		order = append(order, start)

		for _, neighbor := range graph.GetNeighbors(start) {
			if _, ok := visited[neighbor]; !ok {
				order = append(order, Dfs[T](graph, neighbor, visited)...)
			}
		}
	}
	return order
}
