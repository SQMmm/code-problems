package insertionsort

import (
	"fmt"
)

func TestInsertion2() {
	insertionSort2(7, []int32{3, 4, 7, 5, 6, 2, 1})
	fmt.Println()
	insertionSort2(6, []int32{1, 4, 3, 5, 6, 2})
}

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/insertionsort2/problem
 */


func insertionSort2(n int32, arr []int32) {
	fn := func (arr ar, r int32) {
		n := int32(arr[r])
		for i := r-1; i >= 0; i -- {
			if n < arr[i] {
				arr[i+1] = arr[i]
			}else {
				arr[i+1] = n
				fmt.Println(arr)
				return
			}
		}
		arr[0] = n
		fmt.Println(arr)
	}

	for i:= int32(1); i < n; i++ {
		a := ar(arr)
		fn(a, i)
	}
	// Write your code here

}