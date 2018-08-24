package arrays

import (
	"github.com/sqmmm/code-problems/interviewbit/arrays/sort"
	)

func Solve(A []int) int {
	A = sort.SortInt(A)

	if A[len(A) - 1] == 0 {
		return 1
	}

	for i:= 0; i < len(A); i++ {
		index := i
		exit := false
		for A[i] == A[index]{
			if len(A) <= index+1 {
				exit = true
				break
			}
			index++
		}
		if exit {
			continue
		}
		if len(A[index:]) == A[i] {
			return 1
		}
	}

	return -1

}
