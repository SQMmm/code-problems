package treelvtraversal

import "fmt"

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/tree-level-order-traversal/problem
 */

/**
Trees traversal types:
1. Depth First Search or DFS
	1. Inorder Traversal (left node, root, right node)
	2. Preorder Traversal (root, left node, right node)
	3. Postorder Traversal (left node, right node, root)
2. Level Order Traversal or Breadth First Search or BFS
3. Boundary Traversal
4. Diagonal Traversal
 */

type bst struct{
	root *node
}
type node struct {
	value int
	left *node
	right *node
}

func (b *bst) Insert(val int) {
	b.root = b.InsertVal(b.root, val)
}

func (b *bst) InsertVal(n *node, val int) *node {
	if b.root == nil {
		b.root = &node{value: val}
		return b.root
	}

	if n == nil {
		return &node{value: val}
	}

	if val == n.value {
		return n
	}
	if val < n.value {
		n.left = b.InsertVal(n.left, val)
		return n
	}
	n.right = b.InsertVal(n.right, val)
	return n
}

func levelOrder(n *node) {
	queue := []node{*n}
	for i := 0; i < len(queue); i++ {
		fmt.Printf("%v ", queue[i].value)
		if queue[i].left != nil {
			queue = append(queue, *queue[i].left)
		}
		if queue[i].right != nil {
			queue = append(queue, *queue[i].right)
		}
	}
}

func TestLevelOrder() {
	t := &bst{}
	t.Insert(1)
	t.Insert(2)
	t.Insert(5)
	t.Insert(3)
	t.Insert(6)
	t.Insert(4)
	levelOrder(t.root)
}
