package trees

import (
	"reflect"
	"testing"
)

func getTestBinaryTree() *BinaryTreeNode[int] {
	root := NewBinaryTreeNode(5)
	root.Left = NewBinaryTreeNode(3)
	root.Right = NewBinaryTreeNode(7)
	root.Left.Left = NewBinaryTreeNode(2)
	root.Left.Right = NewBinaryTreeNode(4)
	root.Right.Left = NewBinaryTreeNode(6)
	root.Right.Right = NewBinaryTreeNode(8)
	root.Left.Left.Left = NewBinaryTreeNode(1)

	return root
}

func TestPreOrderTraversal(t *testing.T) {
	result := PreOrderTraversal(getTestBinaryTree())
	expected := []int{5, 3, 2, 1, 4, 7, 6, 8}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PreOrderTraversal() = %+v, want %+v", result, expected)
	}
}

func TestInOrderTraversal(t *testing.T) {
	result := InOrderTraversal(getTestBinaryTree())
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InOrderTraversal() = %+v, want %+v", result, expected)
	}
}

func TestPostOrderTraversal(t *testing.T) {
	result := PostOrderTraversal(getTestBinaryTree())
	expected := []int{1, 2, 4, 3, 6, 8, 7, 5}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PostOrderTraversal() = %+v, want %+v", result, expected)
	}
}
