package hightofthebst

import "fmt"

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/tree-height-of-a-binary-tree/problem
 */

func TestGetHeight() {
	t := &BST{}
	t.Insert(3)
	t.Insert(5)
	t.Insert(2)
	t.Insert(1)
	t.Insert(1)
	t.Insert(6)
	t.Insert(7)
	fmt.Println(getHeight(t), "=", 3)
	t = &BST{}
	t.Insert(15)
	fmt.Println(getHeight(t), "=", 0)
	t = &BST{}
	t.Insert(3)
	t.Insert(1)
	t.Insert(7)
	t.Insert(5)
	t.Insert(4)
	fmt.Println(getHeight(t), "=", 3)
}



type node struct {
	value int
	left *node
	right *node
}

type BST struct {
	root *node
}

func (b *BST) Insert(num int) {
	b.InsertVal(b.root, num)
}

func (b *BST) InsertVal(n *node, num int) *node {
	if b.root == nil {
		b.root = &node{value: num}
		return b.root
	}
	if n == nil {
		return &node{value: num}
	}

	if num < n.value {
		n.left  = b.InsertVal(n.left, num)
		return n
	}

	n.right  = b.InsertVal(n.right, num)
	return n
}

func getHeight(t *BST) int {
	return getNodeHeight(t.root, -1)
}

func getNodeHeight(n *node, c int) int {
	if n == nil {
		return c
	}
	c++

	c1 := getNodeHeight(n.left, c)
	c2 := getNodeHeight(n.right, c)

	if c1 > c2 {
		return c1
	}
	return c2
}

