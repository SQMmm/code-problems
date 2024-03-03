package lilyshomework

func LilysHomework(arr []int32) int32 {
	return lilysHomework(arr)
}

/*
 * Complexity: Medium
 * Link: https://www.hackerrank.com/challenges/lilys-homework/problem
 */

func lilysHomework(arr []int32) int32 {
	// Write your code here
	indexMap := make(map[int32]int, len(arr))
	for i, val := range arr {
		indexMap[val] = i
	}

	newArr := make([]int32, len(arr))
	newSLArr := make([]int32, len(arr))
	newSRArr := make([]int32, len(arr))
	copy(newArr, arr)
	copy(newSLArr, arr)
	copy(newSRArr, arr)

	s2 := leftSwapsTmp(arr, mergeSortLeft(newSLArr), getIndexesMap(arr))
	s1 := rightSwapsTmp(newArr, mergeSortRight(newSRArr), getIndexesMap(newArr))


	if s1 < s2 {
		return s1
	}

	return s2
}

func getIndexesMap(a []int32) map[int32]int {
	res := make(map[int32]int, len(a))
	for i, num := range a {
		res[num] = i
	}
	return res
}

func rightSwapsTmp(a, sa []int32, ma map[int32]int) int32 {
	var swaps int32
	for i:= 0; i<len(a)-1; i++ {
		if a[i] == sa[i] {
			continue
		}
		n := a[i] // change index in map
		a[i] = sa[i]
		a[ma[sa[i]]] = n
		ma[n] = ma[sa[i]]
		swaps++
	}


	return swaps
}

func leftSwapsTmp(a, sa []int32, ma map[int32]int) int32  {
	var swaps int32
	for i:= len(a) - 1; i>0; i-- {
		if a[i] == sa[i] {
			continue
		}
		n := a[i] // change index in map
		a[i] = sa[i]
		a[ma[sa[i]]] = n
		ma[n] = ma[sa[i]]
		swaps++
	}

	return swaps
}

// [5, 4, 3, 2, 1]
func mergeSortLeft(arr []int32) []int32 {
	fn  := func (a1, a2 []int32) []int32 {
		res := make([]int32, 0, len(a1) + len(a2))

		i, j := 0, 0
		for i < len(a1) && j < len(a2) {
			if a1[i] > a2[j] {
				res = append(res, a1[i])
				i++
				continue
			}
			res = append(res, a2[j])
			j++
		}

		for ;i < len(a1); i++ {
			res = append(res, a1[i])
		}
		for ;j < len(a2); j++ {
			res = append(res, a2[j])
		}

		return res
	}
	if len(arr) <= 1 {
		return arr
	}

	m := (len(arr))/2
	return fn(mergeSortLeft(arr[:m]), mergeSortLeft(arr[m:]))
}

// [1, 2, 3, 4, 5]
func mergeSortRight(arr []int32) []int32 {
	fn := func (a1, a2 []int32) []int32 {
		res := make([]int32, 0, len(a1) + len(a2))

		i, j := 0, 0
		for i < len(a1) && j < len(a2) {
			if a1[i] < a2[j] {
				res = append(res, a1[i])
				i++
				continue
			}
			res = append(res, a2[j])
			j++
		}

		for ;i < len(a1); i++ {
			res = append(res, a1[i])
		}
		for ;j < len(a2); j++ {
			res = append(res, a2[j])
		}

		return res
	}
	if len(arr) <= 1 {
		return arr
	}

	m := (len(arr))/2
	return fn(mergeSortRight(arr[:m]), mergeSortRight(arr[m:]))
}