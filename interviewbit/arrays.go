package main

import (
	"fmt"
	"github.com/sqmmm/code-problems/interviewbit/arrays"
)

func main() {
	//test min-steps-in-infinite-grid
	fmt.Println(arrays.CoverPoints([]int{0, 1, 1}, []int{0, 1, 2}))
	fmt.Println(arrays.CoverPoints([]int{0, 1, 3}, []int{0, 2, 2}))
}
