package main

import (
	"fmt"
	"github.com/sqmmm/code-problems/interviewbit/arrays"
)

func main() {
	//test wave-array
	fmt.Println(arrays.Wave([]int{4, 1, 3, 2}))
	fmt.Println(arrays.Wave([]int{1, 4, 5, 6, 1, 3, 6}))
	fmt.Println(arrays.Wave([]int{5, 3, 8, 9, 4, 2}))
	fmt.Println(arrays.Wave([]int{5, 3, 8, 9, 4}))
}