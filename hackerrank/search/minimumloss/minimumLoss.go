package minimumloss

/*
 * Complexity: Medium
 * Link: https://www.hackerrank.com/challenges//minimum-loss/problem
 */

func minimumLoss(price []int64) int32 {
	// Write your code here
	sp := quickSort(price)
	mp := getIndexMap(price)

	min := int64(-1)
	for i := 0;i < len(price) - 1; i++ {
		diff := sp[i] - sp[i+1]
		if (diff < min || min == -1) && mp[sp[i]] < mp[sp[i+1]] {
			min = diff
		}
	}

	return int32(min)
}

func getIndexMap(arr []int64) map[int64]int {
	res := make(map[int64]int, len(arr))
	for i, num:= range arr {
		res[num] = i
	}

	return res
}

func quickSort(arr []int64) []int64 {
	if len(arr) <= 1 {
		return arr
	}
	val := arr[0]
	l, r, e := make([]int64, 0), make([]int64, 0), make([]int64, 0)

	for _, num := range arr {
		if num < val {
			r = append(r, num)
			continue
		}
		if num > val {
			l = append(l, num)
			continue
		}
		e = append(e, num)
	}

	return append(quickSort(l), append(e, quickSort(r)...)...)
}