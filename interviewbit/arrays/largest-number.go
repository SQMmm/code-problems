package arrays

import (
	"strings"
	"strconv"
)

func LargestNumber(A []int) (string) {
	tmp := sortInts(A)

	result := make([]string, 0)
	for _, num := range tmp {
		if num > 0 {
			result = append(result, strconv.Itoa(num))
		}
	}
	if len(result) == 0 {
		return "0"
	}

	return strings.Join(result, "")
}

func sortInts(a []int) []int {
	if len(a) == 1 {

		return a
	}

	diff := len(a)/2

	m := merge(sortInts(a[:diff]), sortInts(a[diff:]))
	return m
}



func comp(f, s string) bool {

	n1, _ := strconv.Atoi(f + s)
	n2, _ := strconv.Atoi(s + f)

	if n1 > n2 {
		return true
	}

	return false
}

func merge(f, s []int) []int {
	result := make([]int, 0)
	var i, j int

	for i < len(f) && j < len(s) {
		if comp(strconv.Itoa(f[i]), strconv.Itoa(s[j])) {
			result = append(result, f[i])
			i++
		}else{
			result = append(result, s[j])
			j++
		}
	}

	if i < len(f) {
		for item := i; item < len(f); item++ {
			result = append(result, f[item])
		}
	}else if j < len(s) {
		for item := j; item < len(s); item++ {
			result = append(result, s[item])
		}
	}

	return result
}





