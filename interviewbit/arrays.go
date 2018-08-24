package main

import (
	"fmt"
	"github.com/sqmmm/code-problems/interviewbit/arrays"
)

func main() {
	//test noble-integers
	fmt.Println(arrays.Solve([]int{6, 7, 5})) //-1
	fmt.Println(arrays.Solve([]int{6, 1})) //1
	fmt.Println(arrays.Solve([]int{-1, -2, 0, 0, 0, -3}))
	fmt.Println(arrays.Solve([]int{-10, 1, -6, -2}))
}
