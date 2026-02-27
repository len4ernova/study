package main

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// O(n^2)
func addTwoNumbers1(l1 *Node, l2 *Node) *ListNode {
	n1, n2 := "", ""
	current := l1
	for current != nil {
		n1 += strconv.Itoa(current.data)
		current = current.next
	}

	current = l2
	for current != nil {
		n2 += strconv.Itoa(current.data)
		current = current.next
	}

	r1, err := strconv.Atoi(n1)
	if err != nil {
		return nil
	}
	r2, err := strconv.Atoi(n2)
	if err != nil {
		return nil
	}

	//sum
	r := r1 + r2

	var l *Node

	nr := strconv.Itoa(r)
	list := strings.Split(nr, "")
	for i := 0; i < len(list); i++ {

		d, err := strconv.Atoi(list[i])
		if err != nil {
			return nil
		}

		newNode := &Node{data: d, next: nil}

		if l == nil {
			l = newNode
		} else {
			current = l
			for current.next != nil {
				current = current.next
			}
			current.next = newNode
		}

	}
	var n string
	current = l
	for current != nil {
		n += strconv.Itoa(current.data)
		current = current.next
	}

	return nil
}

func addTwoNumbers2(l1 *Node, l2 *Node) *Node {
	printNodes(l1)
	printNodes(l2)
	resHead := &Node{}
	current := resHead
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := 0
		sum += carry
		if l1 != nil {
			sum += l1.data
			l1 = l1.next
		}
		if l2 != nil {
			sum += l2.data
			l2 = l2.next
		}

		carry = sum / 10
		current.next = &Node{data: sum % 10}
		current = current.next

	}
	printNodes(resHead.next)
	return resHead.next
}
