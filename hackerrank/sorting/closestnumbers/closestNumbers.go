package closestnumbers

import "fmt"

func TestClosestNumbers() {
	fmt.Println(closestNumbers([]int32{5, 2, 3, 4, 1}))
	fmt.Println(closestNumbers([]int32{-20, -3916237, -357920 ,-3620601, 7374819, -7330761, 30, 6246457, -6461594 ,266854 }))
}
/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/closest-numbers/problem
 */

func closestNumbers(arr []int32) []int32 {
	aSorted := mergeSort(arr)
	minDist := int32(10000001)

	for i, num := range aSorted {
		if i == 0 {
			continue
		}
		sub := num - aSorted[i-1]
		if sub < minDist {
			minDist = sub
		}
	}

	res := make([]int32, 0)
	for i, num := range aSorted {
		if i == 0 {
			continue
		}
		if num - aSorted[i-1] == minDist {
			res = append(res, aSorted[i-1], num)
		}
	}

	return res
}

func mergeSort(arr []int32) []int32 {
	if len(arr) == 1 {
		return arr
	}
	if len(arr) == 2 {
		if arr[0] > arr[1]{
			return []int32{arr[1], arr[0]}
		}
		return []int32{arr[0], arr[1]}
	}
	m := len(arr)/2
	return merge(mergeSort(arr[:m]), mergeSort(arr[m:]))
}

func merge(l, r []int32) []int32 {
	res := make([]int32, 0, len(l) + len(r))
	var (
		li, ri int
	)

	for ;li < len(l) && ri < len(r); {
		if l[li] < r[ri] {
			res = append(res, l[li])
			li ++
		}else{
			res = append(res, r[ri])
			ri++
		}
	}

	for li < len(l) {
		res = append(res, l[li])
		li++
	}
	for ri < len(r) {
		res = append(res, r[ri])
		ri++
	}


	return res
}
