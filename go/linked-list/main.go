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
	fmt.Println("Исходный связанный список")
	list.print()
	fmt.Println("Удалим 1ый и последний эл-т:")
	list.DeleteFirstNode()
	list.DeleteLast()
	list.print()
	list.InsertAfterValue(5, 55)
	list.InsertBeforeValue(55, 555)
	list.print()
	fmt.Println("Удалить узел после 3")
	list.deleteAfterValue(3)
	list.print()
	fmt.Println("Удалить узел перед 7")
	list.deleteBeforeValue(7)
	list.print()
	fmt.Println("Найти узел 5")
	cnt := 5
	n := list.findNodeAt(cnt)
	if n != nil {
		fmt.Printf("Значение %d в узле %d\n", n.data, cnt)
	}

	fmt.Printf("Кол-во узлов: %d\n", list.countNodes())

	var l1, l2 LinkedList
	for i := 1; i < 3; i++ {
		l1.InsertAtBack(i)
		l2.InsertAtFront(i * 2)
	}

	addTwoNumbers2(l1.head, l2.head)
}
