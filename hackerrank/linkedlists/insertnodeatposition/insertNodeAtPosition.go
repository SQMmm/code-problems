package insertnodeatposition

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list/problem
 */

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	// Write your code here

	return insert(llist, data, position, 0)
}

func insert(llist *SinglyLinkedListNode, data, position, currPos int32 )*SinglyLinkedListNode {
	if llist == nil {
		return &SinglyLinkedListNode{data: data}
	}

	if currPos == position {
		return &SinglyLinkedListNode{data: data, next: llist}
	}

	llist.next = insert(llist.next, data, position, currPos+1)
	return llist
}