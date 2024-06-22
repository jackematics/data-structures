package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeek(t *testing.T) {
	queue := Queue{
		Length: 0,
	}

	queue.Enqueue(1)

	assert.Equal(t, 1, queue.Peek())
}

func TestDequeue(t *testing.T) {
	queue := Queue{
		Length: 0,
	}

	queue.Enqueue(1)

	assert.Equal(t, 1, queue.Dequeue())
	assert.Equal(t, 0, queue.Length)
}

func TestDequeueNoHead(t *testing.T) {
	queue := Queue{
		Length: 0,
	}

	assert.Equal(t, nil, queue.Dequeue())
}

func TestDoubleDequeue(t *testing.T) {
	queue := Queue{
		Length: 0,
	}

	queue.Enqueue(1)
	queue.Enqueue(2)

	queue.Dequeue()
	assert.Equal(t, 2, queue.Dequeue())
	assert.Equal(t, 0, queue.Length)
}

func TestEnqueueEmpty(t *testing.T) {
	queue := Queue{
		Length: 0,
	}

	queue.Enqueue(1)

	assert.Equal(t, 1, queue.Peek())
	assert.Equal(t, 1, queue.Length)
}

func TestEnqueueTwice(t *testing.T) {
	queue := Queue{
		Length: 0,
	}

	queue.Enqueue(1)
	queue.Enqueue(2)

	assert.Equal(t, 1, queue.Peek())
	assert.Equal(t, 2, queue.Length)
	assert.Equal(t, 2, queue.head.next.value)
}
