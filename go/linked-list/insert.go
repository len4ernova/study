package main

import "fmt"

// insert
func (list *LinkedList) InsertAtBack(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = list.head
	}
	current.next = newNode
}

func (list *LinkedList) InsertAtFront(data int) {
	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
	}

	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func (list *LinkedList) InsertAfterValue(afterValue, data int) {
	newNode := &Node{data: data, next: nil}

	current := list.head

	for current.next != nil {
		if current.data == afterValue {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}

	fmt.Printf("Can't insert node, after value %d doesn't exist", afterValue)
	fmt.Println()
}

func (list LinkedList) InsertBeforeValue(beforeValue, data int) {
	if list.head.next == nil {
		return
	}

	if list.head.data == beforeValue {
		newNode := &Node{data: data, next: list.head}
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		if current.next.data == beforeValue {
			newNode := &Node{data: data, next: current.next}
			current.next = newNode
			return
		}
		current = current.next

	}
	fmt.Printf("Can't insert node, after value %d doesn't exist", beforeValue)
	fmt.Println()
}
