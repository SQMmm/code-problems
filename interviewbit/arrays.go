package main

import (
	"fmt"
	"github.com/sqmmm/code-problems/interviewbit/arrays"
)

func main() {
	//test max-non-negative-subarray
	fmt.Println(arrays.Maxset([]int{1, 2, 5, -7, 2, 3}))
	fmt.Println(arrays.Maxset([]int{1, 2, 3, -7, 1, 5}))
	fmt.Println(arrays.Maxset([]int{1, 2, 3, -7, 3, 1, 2}))
	fmt.Println(arrays.Maxset([]int{-63216, -75973, -82720, -81169, 15827, -16728, -20384, 89949, 16046, -68666, 60520,
	41991, -33258, 68744, 85718, -66825, 58689, -73014, 2918, 33029, 6727, 33754, -8396, 3083, -46865, 15252, -77876, -32675, -82710, 28192, -54747}))

}