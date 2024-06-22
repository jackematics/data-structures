package doubly_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	list := DoublyLinkedList{
		Length: 0,
	}

	for i := range 4 {
		list.Append(i)
	}

	assert.Equal(t, 0, list.Get(0))
	assert.Equal(t, 1, list.Get(1))
	assert.Equal(t, 2, list.Get(2))
	assert.Equal(t, 3, list.Get(3))
}

func TestAppendFirst(t *testing.T) {
	list := DoublyLinkedList{
		Length: 0,
	}

	list.Append(1)

	assert.Equal(t, 1, list.Get(0))
}

func TestAppendTwice(t *testing.T) {
	list := DoublyLinkedList{
		Length: 0,
	}

	list.Append(1)
	list.Append(2)

	assert.Equal(t, 2, list.Get(1))
}

func TestAppendThrice(t *testing.T) {
	list := DoublyLinkedList{
		Length: 0,
	}

	list.Append(1)
	list.Append(2)
	list.Append(5)

	assert.Equal(t, 5, list.Get(2))
}

func TestInsertAt(t *testing.T) {
	list := DoublyLinkedList{
		Length: 0,
	}

	for i := range 4 {
		list.Append(i)
	}

	list.InsertAt(9, 2)

	assert.Equal(t, 1, list.Get(1))
	assert.Equal(t, 9, list.Get(2))
	assert.Equal(t, 2, list.Get(3))
}

func TestInsertAtNegativeIndex(t *testing.T) {
	list := DoublyLinkedList{
		Length: 0,
	}

	for i := range 4 {
		list.Append(i)
	}

	err := list.InsertAt(9, -1)

	assert.Equal(t, err.Error(), "index out of range")
}

func TestInsertAtOutOfRange(t *testing.T) {
	list := DoublyLinkedList{
		Length: 0,
	}

	for i := range 4 {
		list.Append(i)
	}

	err := list.InsertAt(9, 5)

	assert.Equal(t, err.Error(), "index out of range")
}

func TestInsertAtFirstElement(t *testing.T) {
	list := DoublyLinkedList{
		Length: 0,
	}

	list.InsertAt(0, 0)

	assert.Equal(t, 0, list.Get(0))
}

func TestInsertAtLastElement(t *testing.T) {
	list := DoublyLinkedList{
		Length: 0,
	}

	list.Append(0)

	list.InsertAt(1, 0)

	assert.Equal(t, 0, list.Get(1))
	assert.Equal(t, 1, list.Get(0))
	assert.Equal(t, 1, list.Head.value)
}
