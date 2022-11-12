package structures

type Stack[T any] struct {
	index    uint64
	capacity uint64
	values   []T
}

func (stack *Stack[T]) Push(value T) {
	if stack.index >= stack.capacity {
		stack.values = append(stack.values, value)
		stack.capacity++
	} else {
		stack.values[stack.index] = value
	}

	stack.index++
}

func (stack *Stack[T]) Pop() T {
	stack.index--

	top := stack.values[stack.index]

	return top
}

func (stack *Stack[T]) Top() T {
	return stack.values[stack.index-1]
}
