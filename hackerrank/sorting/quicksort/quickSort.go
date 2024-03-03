package quicksort

import "fmt"

func TestQuickSortPartition() {
	fmt.Println(quickSortPartition([]int32{5, 7, 4, 3, 8}))
	fmt.Println(quickSortPartition([]int32{4, 5, 3, 7, 2}))
}

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/quicksort1/problem
 */

func quickSortPartition(arr []int32) []int32 {
	p := arr[0]
	l, r, e := make([]int32, 0), make([]int32, 0), []int32{p}

	for i := range arr  {
		if i == 0 {
			continue
		}
		if arr[i] < p {
			l = append(l, arr[i])
		}else if arr[i] > p {
			r  = append(r, arr[i])
		}else{
			e = append(e, arr[i])
		}
	}

	return append(l, append(e, r...)...)

}