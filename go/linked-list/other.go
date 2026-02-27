package main

import "fmt"

func (list *LinkedList) countNodes() int {
	var count int
	current := list.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// нахождение узла по индексу
func (list *LinkedList) findNodeAt(index int) *Node {
	count := 0
	current := list.head

	for current != nil {
		count++
		current = current.next
	}

	if index < 0 || index > count {
		return nil
	}

	current = list.head
	for i := 1; i < index; i++ {
		current = current.next
	}
	return current
}
func (list *LinkedList) print() {
	var current *Node = list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println()
}

func printNodes(l *Node) {
	var current *Node = l
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println()
}
