package trees

type BinaryTreeNode[T any] struct {
	Value T
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}

func NewBinaryTreeNode[T any](value T) *BinaryTreeNode[T] {
	return &BinaryTreeNode[T]{
		Value: value,
	}
}

func PreOrderTraversal[T any](node *BinaryTreeNode[T]) []T {
	result := make([]T, 0)

	if node == nil {
		return result
	}

	result = append(result, node.Value)
	result = append(result, PreOrderTraversal[T](node.Left)...)
	result = append(result, PreOrderTraversal[T](node.Right)...)

	return result
}

func InOrderTraversal[T any](node *BinaryTreeNode[T]) []T {
	result := make([]T, 0)

	if node == nil {
		return result
	}

	result = append(result, InOrderTraversal[T](node.Left)...)
	result = append(result, node.Value)
	result = append(result, InOrderTraversal[T](node.Right)...)

	return result
}

func PostOrderTraversal[T any](node *BinaryTreeNode[T]) []T {
	result := make([]T, 0)

	if node == nil {
		return result
	}

	result = append(result, PostOrderTraversal[T](node.Left)...)
	result = append(result, PostOrderTraversal[T](node.Right)...)
	result = append(result, node.Value)

	return result
}
