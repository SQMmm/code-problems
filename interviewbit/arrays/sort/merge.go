package sort

func SortInt(A []int) []int {
	if len(A) <= 1 {
		return A
	}

	mid := len(A)/2

	left, right := split(A, mid)
	left = SortInt(left)
	right = SortInt(right)
	return merge(left, right)
}

func split(A []int, mid int) ([]int, []int)  {
	return A[:mid], A[mid:]
}

func merge(A, B []int) []int {
	arr := make([]int, len(A) + len(B))

	// index j for A, k for B
	j, k := 0, 0

	for i := 0; i < len(arr); i++ {
		// fix for index out of range without using sentinel
		if j >= len(A) {
			arr[i] = B[k]
			k++
			continue
		} else if k >= len(B) {
			arr[i] = A[j]
			j++
			continue
		}
		// default loop condition
		if A[j] > B[k] {
			arr[i] = B[k]
			k++
		} else {
			arr[i] = A[j]
			j++
		}
	}

	return arr
}

