package arrays

import "math"

func CoverPoints(A []int , B []int )  (int) {
	count := 0
	for i := range A {
		if i == 0 {
			continue
		}

		count += max(int(math.Abs(float64(A[i] - A[i-1]))), int(math.Abs(float64(B[i] - B[i-1]))))
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}