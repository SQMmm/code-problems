package reverselist


/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/reverse-a-linked-list/problem
 */

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func reverse(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	// Write your code here

	if llist == nil {
		return nil
	}

	return reverselist(llist.next, &SinglyLinkedListNode{data:llist.data})
}

func reverselist(initL, newL *SinglyLinkedListNode) *SinglyLinkedListNode {
	if initL == nil {
		return newL
	}

	newL = &SinglyLinkedListNode{data: initL.data, next: newL}
	return reverselist(initL.next, newL)

}