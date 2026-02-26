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

	current := list.head
	for current.next != nil {

		current = current.next
	}
	current.next = nil
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
	fmt.Printf("Node with after value %d  has data %d and will be deleted", afterValue, tmp.data)
	current.next = current.next.next

}

// TODO
func (list *LinkedList) deleteBeforeValue(deforeValue int) {

}
