package interviewbit


func Zipper(a, b []int, n int) []int {
	res := make([]int, 0, n)
	if len(a) < n {
		n = len(a)
	}
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		res = append(res, a[i], b[i])
	}

	return res
}


