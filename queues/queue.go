package queues

import "fmt"

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

func (q *Queue[T]) Push(item T) *Queue[T] {
	q.items = append(q.items, item)
	return q
}

func (q *Queue[T]) PopLeft() (T, error) {
	if q.Empty() {
		return *new(T), fmt.Errorf("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue[T]) PopRight() (T, error) {
	if q.Empty() {
		return *new(T), fmt.Errorf("queue is empty")
	}

	item := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return item, nil
}

func (q *Queue[T]) Empty() bool {
	return len(q.items) == 0
}
