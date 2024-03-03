package maxtriangleperimeter

import "fmt"

func TestMaximumPerimeterTriangle() {
	fmt.Println(maximumPerimeterTriangle([]int32{8, 1, 2, 3, 4, 5, 10}), "=", "3 4 5")
	//fmt.Println(maximumPerimeterTriangle([]int32{1, 1, 1, 3, 3}), "=", "1 3 3")
	//fmt.Println(maximumPerimeterTriangle([]int32{1, 2, 3}), "=", "-1")
	//fmt.Println(maximumPerimeterTriangle([]int32{1, 1, 1, 2, 3, 5}), "=", "1 1 1")
}

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/maximum-perimeter-triangle/problem
 */

func maximumPerimeterTriangle(sticks []int32) []int32 {
	sticks = mergeSort(sticks, 0, len(sticks) - 1)


	for i := len(sticks) - 1; i >= 2; i-- {
		a, b, c := sticks[i], sticks[i - 1], sticks[i - 2]
		if (b + c) > a {
			return []int32{c, b, a}
		}
	}

	return []int32{-1}
}

func mergeSort(arr []int32, l, r int) []int32 {
	if l == r {
		return []int32{arr[l]}
	}

	m := (r - l) / 2 + l
	return merge(mergeSort(arr, l, m), mergeSort(arr, m+1, r))
}

func merge(a, b []int32) []int32 {
	i, j := 0, 0

	res := make([]int32, 0, len(a) + len(b))
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
			continue
		}
		res = append(res, b[j])
		j++
	}

	for ; i < len(a); i++ {
		res = append(res, a[i])
	}
	for ; j < len(b); j++ {
		res = append(res, b[j])
	}

	return res
}


