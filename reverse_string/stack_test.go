package reverse_string

import "testing"

func TestPushAddsItemToStack(t *testing.T) {
	stack := NewStack[int]()
	item := 42

	stack.Push(&item)

	if len(stack.items) != 1 || *stack.items[0] != item {
		t.Errorf("Push() did not add the item to the stack correctly")
	}
}

func TestPopRemovesAndReturnsLastItem(t *testing.T) {
	stack := NewStack[int]()
	item1, item2 := 42, 84
	stack.Push(&item1).Push(&item2)

	popped := stack.Pop()

	if popped == nil || *popped != item2 {
		t.Errorf("Pop() = %v, want %v", popped, item2)
	}
	if len(stack.items) != 1 || *stack.items[0] != item1 {
		t.Errorf("Pop() did not remove the last item correctly")
	}
}

func TestPopReturnsNilWhenStackIsEmpty(t *testing.T) {
	stack := NewStack[int]()

	popped := stack.Pop()

	if popped != nil {
		t.Errorf("Pop() = %v, want nil", popped)
	}
}
