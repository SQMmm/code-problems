package lowestcommonancestorbst

import "fmt"

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/binary-search-tree-lowest-common-ancestor/problem
 */

func TestIca() {
	t := &bst{}
	t.Insert(4)
	t.Insert(2)
	t.Insert(3)
	t.Insert(1)
	t.Insert(7)
	t.Insert(6)
	fmt.Println(Ica(t.root, 1, 7), "=", 4)
	t = &bst{}
	t.Insert(1)
	t.Insert(2)
	fmt.Println(Ica(t.root, 1, 2), "=", 1)
	t = &bst{}
	t.Insert(5)
	t.Insert(3)
	t.Insert(8)
	t.Insert(2)
	t.Insert(4)
	t.Insert(6)
	t.Insert(7)
	fmt.Println(Ica(t.root, 7, 3), "=", 5)
}

type node struct {
	value int
	left *node
	right *node
}

type bst struct {
	root *node
}

func (b *bst) Insert(num int) {
	b.InsertVal(b.root, num)
}

func (b *bst) InsertVal(n *node, num int) *node {
	if b.root == nil {
		b.root = &node{value: num}
	}

	if n == nil {
		return &node{value: num}
	}

	if num < n.value {
		n.left = b.InsertVal(n.left, num)
		return n
	}
	n.right = b.InsertVal(n.right, num)
	return n
}

func Ica(n *node, v1, v2 int) int {
	if (v1 >= n.value && v2 <= n.value) || (v2 >= n.value && v1 <= n.value) {
		return n.value
	}
	if v1 < n.value {
		return Ica(n.left, v1, v2)
	}
	return Ica(n.right, v1, v2)
}