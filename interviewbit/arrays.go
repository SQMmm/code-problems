package main

import (
	"fmt"
	"github.com/sqmmm/code-problems/interviewbit/arrays"
)

func main() {
	//test min-steps-in-infinite-grid
	fmt.Println(arrays.MaxSpecialProduct([]int{5, 9, 6, 8, 6, 4, 6, 9, 5, 4, 9 }))
	fmt.Println(arrays.MaxSpecialProduct([]int{6, 7, 9, 5, 5, 8 }))
	fmt.Println(arrays.MaxSpecialProduct([]int{7, 5, 7, 9, 8, 7}))
}
