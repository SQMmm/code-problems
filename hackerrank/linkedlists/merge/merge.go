package merge

import "fmt"

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/merge-two-sorted-linked-lists/problem
 */



type node struct {
	data int
	next *node
}

func (n *node) InsertVal(val int) *node {
	if n == nil {
		n = &node{data: val}
		return n
	}

	n.next = n.next.InsertVal(val)
	return n
}

func (n *node) Print() {
	if n == nil {
		fmt.Println()
		return
	}
	fmt.Printf("%v ",n.data)
	n.next.Print()
}

func mergeLists(fl, sl *node) *node {
	return merge(fl, sl, nil)
}

func merge(fl, sl, newl *node) *node {
	if fl == nil && sl == nil {
		return newl
	}

	if sl == nil || fl != nil && fl.data < sl.data {
		newl = &node{data: fl.data}
		newl.next = merge(fl.next, sl, newl.next)
		return newl
	}

	newl = &node{data: sl.data}
	newl.next = merge(fl, sl.next, newl.next)
	return newl
}