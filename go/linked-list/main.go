package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {
	var list LinkedList
	for i := 0; i < 10; i++ {

		list.InsertAtBack(i)
	}

	fmt.Println(list.head.data, list.head.next.data)

}
