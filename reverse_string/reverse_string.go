package reverse_string

func Reverse(input string) string {
	stack := NewStack[rune]()

	for _, char := range input {
		stack.Push(&char)
	}

	result := ""

	for char := stack.Pop(); char != nil; char = stack.Pop() {
		result = result + string(*char)
	}

	return result
}
