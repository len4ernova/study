package main

import (
	"fmt"
)

func (list *LinkedList) DeleteFirstNode() {
	if list.head != nil {
		list.head = list.head.next
		fmt.Println("Head node of linked list has been deleted")
		return
	}
	fmt.Println("Linked list is empty")
}

func (list *LinkedList) DeleteLast() {
	if list.head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	if list.head.next == nil {
		list.head = nil
		fmt.Println("Head node of linked list has been deleted")
		return
	}

	previous := list.head
	current := previous.next
	for current.next != nil {
		previous = current
		current = current.next
	}
	previous.next = nil
	fmt.Println("Head node of linked list has been deleted")
}

func (list *LinkedList) deleteAfterValue(afterValue int) {
	current := list.head

	for current != nil && current.data != afterValue {
		current = current.next
	}

	if current == nil {
		fmt.Printf("Node with after value %d doesn't exist \n", afterValue)
		return
	}

	if current.next == nil {
		fmt.Printf("Node with after value %d is the last node\n", afterValue)
		return
	}

	tmp := current.next
	fmt.Printf("Node with after value %d  has data %d and will be deleted\n", afterValue, tmp.data)
	current.next = current.next.next

}

func (list *LinkedList) deleteBeforeValue(beforeValue int) {
	previous := list.head

	if previous == nil || previous.next == nil {
		fmt.Printf("Node before value %d doesn't exist \n", beforeValue)
		return
	}

	current := previous.next
	if current.data == beforeValue {
		list.head = current
		return
	}

	for current.next.data != beforeValue {
		if current.next == nil {
			fmt.Printf("Node with value %d doesn't exist \n", beforeValue)
			return
		}
		previous = current
		current = current.next
	}
	previous.next = current.next
}

func (list *LinkedList) deleteBeforeValue2(beforeValue int) {
	var previous *Node
	current := list.head

	if current == nil || current.next == nil {
		fmt.Printf("Node with before value doesn't exist\n")
		return
	}

	for current.next != nil {
		if current.next.data == beforeValue {
			if previous == nil {
				fmt.Printf("Node before value %d has data %d and will be deleted\n", beforeValue, current.data)
				list.head = current.next
				return
			} else {
				fmt.Printf("Node before value %d has data %d and will be deleted\n", beforeValue, current.data)
				previous.next = current.next
			}
			return
		}
		previous = current
		current = current.next
	}
	fmt.Printf("Node before value %d not found\n", beforeValue)
}
