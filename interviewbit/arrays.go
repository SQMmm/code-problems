package main

import (
	"fmt"
	"github.com/sqmmm/code-problems/interviewbit/arrays"
)

func main() {
	//test largest-namber
	fmt.Println(arrays.LargestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(arrays.LargestNumber([]int{3, 33, 34, 5, 9}))
	fmt.Println(arrays.LargestNumber([]int{3, 32, 34, 5, 9}))
	fmt.Println(arrays.LargestNumber([]int{101, 3, 32, 34, 5, 9}))
	fmt.Println(arrays.LargestNumber([]int{301, 3, 32, 34, 5, 9}))
	fmt.Println(arrays.LargestNumber([]int{29, 298}))
	fmt.Println(arrays.LargestNumber([]int{29, 2987}))
	fmt.Println(arrays.LargestNumber([]int{12, 121}))
	fmt.Println(arrays.LargestNumber([]int{931, 94, 209, 448, 716, 903, 124, 372, 462, 196, 715, 802, 103, 740, 389, 872, 615, 638, 771, 829, 899, 999, 29, 163, 342, 902, 922, 312, 326, 817, 288, 75, 37, 286, 708, 589, 975, 747, 743, 699, 743, 954, 523, 989, 114, 402, 236, 855, 323, 79, 949, 176, 663, 587, 322}))
	fmt.Println(arrays.LargestNumber([]int{0,0,0}))

}