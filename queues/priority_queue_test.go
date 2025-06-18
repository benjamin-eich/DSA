package queues

import (
	"container/heap"
	"testing"
)

func getTestPriorityQueue() (*PriorityQueue[int], []*PriorityQueueItem[int]) {
	items := []*PriorityQueueItem[int]{
		{value: 1, priority: 2},
		{value: 2, priority: 3},
		{value: 3, priority: 1},
	}
	pq := &PriorityQueue[int]{}
	for _, item := range items {
		heap.Push(pq, item)
	}
	return pq, items
}

func TestPriorityQueue_PushAndPop(t *testing.T) {
	pq := &PriorityQueue[int]{}
	heap.Push(pq, &PriorityQueueItem[int]{value: 10, priority: 1})
	heap.Push(pq, &PriorityQueueItem[int]{value: 20, priority: 3})
	heap.Push(pq, &PriorityQueueItem[int]{value: 30, priority: 2})

	wantLen := 3
	if pq.Len() != wantLen {
		t.Errorf("Len() = %d, want %d", pq.Len(), wantLen)
	}

	item := heap.Pop(pq).(*PriorityQueueItem[int])
	wantVal := 20
	if item.value != wantVal {
		t.Errorf("Pop() = %v, want %v", item.value, wantVal)
	}

	wantLen = 2
	if pq.Len() != wantLen {
		t.Errorf("Len() = %d, want %d", pq.Len(), wantLen)
	}
}

func TestPriorityQueue_LessAndSwap(t *testing.T) {
	pq := PriorityQueue[int]{
		{value: 1, priority: 2, index: 0},
		{value: 2, priority: 3, index: 1},
	}
	if !pq.Less(1, 0) {
		t.Errorf("Less(1,0) = %t, want true", pq.Less(1, 0))
	}
	pq.Swap(0, 1)
	if pq[0].value != 2 || pq[1].value != 1 {
		t.Errorf("Swap() = %v, want [2 1]", []int{pq[0].value, pq[1].value})
	}
	if pq[0].index != 0 || pq[1].index != 1 {
		t.Errorf("Swap() index = [%d %d], want [0 1]", pq[0].index, pq[1].index)
	}
}

func TestPriorityQueue_PushTypeAssertion(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Push() did not panic on wrong type, want panic")
		}
	}()
	pq := &PriorityQueue[int]{}
	pq.Push("not-an-item")
}

func TestPriorityQueue_update(t *testing.T) {
	pq, items := getTestPriorityQueue()
	item := items[0]
	pq.update(item, 99, 10)
	if item.value != 99 {
		t.Errorf("update() value = %v, want %v", item.value, 99)
	}
	if item.priority != 10 {
		t.Errorf("update() priority = %v, want %v", item.priority, 10)
	}
	// Nach update sollte das Item an der richtigen Stelle im Heap sein
	max := heap.Pop(pq).(*PriorityQueueItem[int])
	if max.value != 99 {
		t.Errorf("update() max value = %v, want %v", max.value, 99)
	}
}
