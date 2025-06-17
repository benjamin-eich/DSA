package graphs

import "DSA/queues"

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
