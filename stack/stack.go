package stack

type Node struct {
	value any
	next  *Node
}

type Stack struct {
	Length int
	head   *Node
}

func (stack *Stack) Peek() any {
	return stack.head.value
}

func (stack *Stack) Pop() any {
	if stack.head == nil {
		return nil
	}

	val := stack.head.value
	stack.head = stack.head.next
	stack.Length--

	return val
}

func (stack *Stack) Push(val any) {
	stack.Length++
	newNode := &Node{value: val}

	if stack.head == nil {
		stack.head = newNode
		return
	}

	newNode.next = stack.head
	stack.head = newNode
}
