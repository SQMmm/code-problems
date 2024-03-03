package selfbalancingbst

import (
	"fmt"
)

/*
 * Complexity: Medium
 * Link: https://www.hackerrank.com/challenges/self-balancing-tree/problem
 */

func TestAvl() {
	t := NewAvl()
	//t.Insert(3)
	//t.Insert(2)
	//t.Insert(4)
	//t.Insert(5)
	//t.Insert(6)
	//t.Print()
	t = NewAvl()
	t.Insert(14)
	t.Insert(25)
	t.Insert(21)
	//t.Insert(10)
	//t.Insert(23)
	//t.Insert(7)
	//t.Insert(26)
	//t.Insert(12)
	//t.Insert(30)
	//t.Insert(16)
	//t.Insert(19)
	t.Print()
}

type node struct {
	value int
	left *node
	right *node
	height int
}

type avl struct  {
	root *node
}

func NewAvl() *avl {
	return &avl{root: &node{height: -1}}
}

func (a *avl) Insert(num int) {
	a.root, _ = a.InsertVal(a.root, num)
}

func (a *avl) InsertVal(n *node, num int) (*node, int) {
	if a.root.height == -1 {
		a.root = &node{value: num, height: 0, left: &node{height: -1}, right: &node{height: -1}}
		return a.root, 0
	}

	if n.height == -1 {
		return &node{value: num, height: 0, left: &node{height: -1}, right: &node{height: -1}}, 0
	}

	var h int
	if num < n.value {
		n.left, h = a.InsertVal(n.left, num)
	}else {
		n.right, h = a.InsertVal(n.right, num)
	}

	h++
	if h > n.height {
		n.height = h
	}

	bf := n.left.height - n.right.height
	if bf > 1 {
		if num < n.left.value {
			n = rotateRight(n)
			return n, n.height
		}
		n.left = rotateLeft(n.left)
		n = rotateRight(n)
		return n, n.height
	}else if bf < -1 {
		if num > n.right.value{
			n = rotateLeft(n)
			return n, n.height
		}
		n.right = rotateRight(n.right)
		n = rotateLeft(n)
		return n, n.height

	}

	return n, n.height
}

func rotateRight(n *node) *node {
	l := n.left
	l.right, n.left = n, l.right


	n.height = n.maxHeight() + 1
	l.height = l.maxHeight() + 1
	return l
}

func rotateLeft(n *node) *node {
	r := n.right
	r.left, n.right = n, r.left

	n.height = n.maxHeight() + 1
	r.height = r.maxHeight() + 1

	return r
}

func (n *node) maxHeight() int {
	if n.left.height > n.right.height {
		return n.left.height
	}
	return n.right.height
}

func (a *avl) Print() {
	a.root.printOrder()
	fmt.Println()
	a.root.printTree()
}

func (n *node) printOrder() {
	if n.height == -1 {
		return
	}
	if n.height == 0 {
		fmt.Printf("%v(BF=%v) ", n.value, n.left.height - n.right.height)
		return
	}
	n.left.printOrder()
	fmt.Printf("%v(BF=%v) ", n.value, n.left.height - n.right.height)
	n.right.printOrder()
}

func (n *node) printTree() {
	if n.height == -1 {
		return
	}
	fmt.Printf("%v(BF=%v) ", n.value, n.left.height - n.right.height)
	n.left.printTree()
	n.right.printTree()
}