package arrays

import (
	"github.com/sqmmm/code-problems/interviewbit/arrays/sort"
)

func Wave(A []int ) []int {
	if len(A) < 2 {
		return A
	}

	A = sort.SortInt(A)

	for i := 1; i < len(A); i += 2 {
		A[i], A[i - 1] = A[i - 1], A[i]
	}

	return A
}


