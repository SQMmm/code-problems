package deletenode

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/delete-a-node-from-a-linked-list/problem
 */


type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func deleteNode(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	// Write your code here

	return delete(llist, position, 0)
}


func delete(llist *SinglyLinkedListNode, position, currPos int32) *SinglyLinkedListNode {
	if llist == nil {
		return llist
	}

	if currPos == position {
		if llist.next == nil {
			return nil
		}
		return &SinglyLinkedListNode{data: llist.next.data, next: llist.next.next}
	}

	llist.next = delete(llist.next, position, currPos + 1)
	return llist
}