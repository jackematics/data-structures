package doubly_linked_list

import (
	"errors"
	"fmt"
)

type Node struct {
	value any
	next  *Node
	prev  *Node
}

type DoublyLinkedList struct {
	Length int
	Head   *Node
}

func (list *DoublyLinkedList) Get(index int) any {
	if index < 0 || index > list.Length-1 {
		panic("out of range")
	}

	currentNode := list.Head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	return currentNode.value
}

func (list *DoublyLinkedList) Append(item any) {
	newNode := Node{value: item}
	list.Length++

	if list.Head == nil {
		list.Head = &newNode
		return
	}

	currentNode := list.Head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = &newNode
	newNode.prev = currentNode
}

func (list *DoublyLinkedList) InsertAt(item any, index int) error {
	list.Length++
	if index < 0 || index > list.Length-1 {
		list.Length--
		return errors.New("index out of range")
	}

	newNode := &Node{value: item}

	if list.Head == nil {
		list.Head = newNode
		return nil
	}

	currentNode := list.Head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	newNode.next = currentNode

	if currentNode.prev != nil {
		newNode.prev = currentNode.prev
		currentNode.prev.next = newNode
		currentNode.prev = newNode
	} else {
		list.Head = newNode
	}

	return nil
}

func (list *DoublyLinkedList) Print(message string) {
	print(message + ": [")
	currentNode := list.Head
	for currentNode != nil {
		fmt.Printf("%+v", currentNode.value)

		if currentNode.next != nil {
			print(", ")
		}
		currentNode = currentNode.next
	}
	print("]")
}
