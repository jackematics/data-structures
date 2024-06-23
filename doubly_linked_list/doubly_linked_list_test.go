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

func TestRemove(t *testing.T) {
	list := DoublyLinkedList{Length: 0}

	list.Append(0)
	list.Append(1)
	list.Append(2)

	removed := list.Remove(1)

	assert.Equal(t, 2, list.Length)
	assert.Equal(t, 1, removed)
	assert.Equal(t, 0, list.Get(0))
	assert.Equal(t, 2, list.Get(1))
}

func TestRemoveFirst(t *testing.T) {
	list := DoublyLinkedList{Length: 0}

	list.Append(0)

	removed := list.Remove(0)

	nilNode := &Node{}
	nilNode = nil

	assert.Equal(t, 0, list.Length)
	assert.Equal(t, 0, removed)
	assert.Equal(t, nilNode, list.Head)
}

func TestRemoveNonExistent(t *testing.T) {
	list := DoublyLinkedList{Length: 0}

	list.Append(0)

	removed := list.Remove(1)

	assert.Equal(t, 1, list.Length)
	assert.Equal(t, nil, removed)
}

func TestRemoveAtLastItem(t *testing.T) {
	list := DoublyLinkedList{Length: 0}

	list.Append(0)

	val := list.RemoveAt(0)

	nilNode := &Node{}
	nilNode = nil

	assert.Equal(t, 0, val)
	assert.Equal(t, 0, list.Length)
	assert.Equal(t, nilNode, list.Head)
}

func TestRemoveAt(t *testing.T) {
	list := DoublyLinkedList{Length: 0}

	list.Append(0)
	list.Append(1)
	list.Append(2)
	list.Append(3)

	val := list.RemoveAt(2)

	assert.Equal(t, 3, list.Length)
	assert.Equal(t, 2, val)

	assert.Equal(t, 0, list.Get(0))
	assert.Equal(t, 1, list.Get(1))
	assert.Equal(t, 3, list.Get(2))
}

func TestPrependOnEmpty(t *testing.T) {
	list := DoublyLinkedList{Length: 0}

	list.Prepend(0)

	assert.Equal(t, 1, list.Length)
	assert.Equal(t, 0, list.Get(0))
}

func TestPrepend(t *testing.T) {
	list := DoublyLinkedList{Length: 0}

	list.Append(0)
	list.Append(1)

	list.Prepend(9)

	assert.Equal(t, 9, list.Get(0))
}
