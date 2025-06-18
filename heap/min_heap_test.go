package heap

import (
	"container/heap"
	"testing"
)

func TestMinHeap_PushAndPop(t *testing.T) {
	h := &MinHeap{}
	heap.Push(h, 3)
	heap.Push(h, 1)
	heap.Push(h, 2)

	wantLen := 3
	if h.Len() != wantLen {
		t.Errorf("Len() = %d, want %d", h.Len(), wantLen)
	}

	val := heap.Pop(h).(int)
	wantVal := 1
	if val != wantVal {
		t.Errorf("Pop() = %d, want %d", val, wantVal)
	}

	wantLen = 2
	if h.Len() != wantLen {
		t.Errorf("Len() = %d, want %d", h.Len(), wantLen)
	}
}

func TestMinHeap_LessAndSwap(t *testing.T) {
	h := MinHeap{5, 2, 8}
	if !h.Less(1, 0) {
		t.Errorf("Less(1,0) = %t, want true", h.Less(1, 0))
	}
	h.Swap(0, 2)
	want := []int{8, 2, 5}
	for i, v := range want {
		if h[i] != v {
			t.Errorf("Swap() = %v, want %v", h, want)
			break
		}
	}
}

func TestMinHeap_PushTypeAssertion(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Push() did not panic on wrong type, want panic")
		}
	}()
	h := &MinHeap{}
	h.Push("not-an-int")
}
