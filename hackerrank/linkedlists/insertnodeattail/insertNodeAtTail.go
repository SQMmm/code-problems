package insertnodeattail

import "fmt"

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/insert-a-node-at-the-tail-of-a-linked-list/problem
 */

type node struct {
	value int
	next *node
}

type list struct {
	head *node
}

func (l *list) Insert(num int) {
	newNode := &node{value: num}
	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (l *list) Print() {
	if l.head == nil {
		return
	}

	curr := l.head
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func insertNodeAtTail(arr []int) *list {
	l := &list{}
	for _, v := range arr {
		l.Insert(v)
	}

	return l
}



