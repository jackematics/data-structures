package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeek(t *testing.T) {
	stack := Stack{
		Length: 0,
	}

	stack.Push(0)

	assert.Equal(t, 0, stack.Peek())
}

func TestPop(t *testing.T) {
	stack := Stack{
		Length: 0,
	}

	stack.Push(0)

	assert.Equal(t, 0, stack.Pop())
	assert.Equal(t, 0, stack.Length)
}

func TestPopNoHead(t *testing.T) {
	stack := Stack{
		Length: 0,
	}

	assert.Equal(t, nil, stack.Pop())
}

func TestDoublePop(t *testing.T) {
	stack := Stack{
		Length: 0,
	}

	stack.Push(0)
	stack.Push(1)

	stack.Pop()
	assert.Equal(t, 0, stack.Pop())
	assert.Equal(t, 0, stack.Length)
}

func TestPushEmpty(t *testing.T) {
	stack := Stack{
		Length: 0,
	}

	stack.Push(1)

	assert.Equal(t, 1, stack.Peek())
	assert.Equal(t, 1, stack.Length)
}

func TestPushTwice(t *testing.T) {
	stack := Stack{
		Length: 0,
	}

	stack.Push(1)
	stack.Push(2)

	assert.Equal(t, 2, stack.Peek())
	assert.Equal(t, 2, stack.Length)
	assert.Equal(t, 1, stack.head.next.value)
}
