package reverseprint

import "fmt"

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/print-the-elements-of-a-linked-list-in-reverse/problem
 */

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func reversePrint(llist *SinglyLinkedListNode) {
	// Write your code here
	if llist == nil {
		return
	}

	reversePrint(llist.next)
	fmt.Println(llist.data)
}