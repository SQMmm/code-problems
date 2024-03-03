package comparelists

import "fmt"

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/compare-two-linked-lists/problem
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
func (n *node) Print() {
	curr := n
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func compare_lists(l1, l2 *node) int32 {
	for l1 != nil && l2 != nil {
		if l1.value != l2.value {
			return 0
		}
		l1 = l1.next
		l2 = l2.next
	}

	if l1 != nil || l2 != nil {
		return 0
	}

	return 1
}