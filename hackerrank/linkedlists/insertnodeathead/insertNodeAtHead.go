package insertnodeathead

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/insert-a-node-at-the-head-of-a-linked-list/problem
 */


type node struct {
	data int32
	next *node
}

func (n *node) Print() {
	if n.next != nil {
		n.next.Print()
	}
}

func insertNodeAtHead(n *node, data int32) *node {
	return &node{data: data, next: n}
}
