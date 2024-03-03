package maxmin

import "fmt"

/*
 * Complexity: Medium
 * Link: https://www.hackerrank.com/challenges/angry-children/problem
 */

func TestMaxMin() {
	fmt.Println(maxMin(3, []int32{10, 100, 300, 200, 1000, 20, 30}), "=", 20)
	fmt.Println(maxMin(2, []int32{1, 4, 7, 2}), "=", 1)
	fmt.Println(maxMin(4, []int32{1, 2, 3, 4, 10, 20, 30, 40, 100, 200}), "=", 3)
	fmt.Println(maxMin(2, []int32{1, 2, 1, 2, 1}), "=", 0)
}

func maxMin(k int32, arr []int32) int32 {
	// Write your code here
	arr = quickSort(arr)

	min := int64(1000000001)
	for i := int32(0); i <= int32(len(arr)) - k; i++ {
		diff := int64(arr[i+k-1] - arr[i])
		if diff < min {
			min = diff
		}
	}

	return int32(min)
}


func quickSort(arr []int32) []int32 {
	if len(arr) <= 1 {
		return arr
	}

	val := arr[0]
	l, r, e := make([]int32, 0), make([]int32, 0),make([]int32, 0)

	for _, a := range arr {
		if a < val {
			l = append(l, a)
			continue
		}
		if a > val {
			r = append(r, a)
			continue
		}
		e = append(e, a)
	}

	return append(quickSort(l), append(e, quickSort(r)...)...)
}