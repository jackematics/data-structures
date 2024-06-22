package queue

type Node struct {
	value any
	next  *Node
}

type Queue struct {
	Length int
	head   *Node
	tail   *Node
}

func (queue *Queue) Peek() any {
	return queue.head.value
}

func (queue *Queue) Dequeue() any {
	if queue.head == nil {
		return nil
	}

	val := queue.head.value
	queue.head = queue.head.next
	queue.Length--

	return val
}

func (queue *Queue) Enqueue(val any) {
	queue.Length++
	newNode := &Node{value: val}
	if queue.head == nil {
		queue.head = newNode
		queue.tail = queue.head
	}

	queue.tail.next = newNode
	queue.tail = newNode
}
