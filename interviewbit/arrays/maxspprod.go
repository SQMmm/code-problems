package arrays

func MaxSpecialProduct (A []int )  (int)  {
	rv := make([]int, 0)
	ri := make([]int, 0)
	lv := make([]int, 0)
	li := make([]int, 0)

	maxL := make([]int, len(A))
	max := make([]int, len(A))

	for i, v := range A {
		if len(lv) > 0 {
			for j := len(lv) - 1; j >= 0; j-- {
				if lv[j] > v {
					maxL[i] = li[j]
					break
				}
			}
		}

		lv = append(lv, v)
		li = append(li, i)
	}

	for i := len(A) - 1; i >= 0 ; i-- {
		if len(rv) > 0 && maxL[i] > 0 {
			for j := len(rv) - 1; j >= 0; j-- {
				if rv[j] > A[i] {
					max[i] = ri[j] * maxL[i]
					break
				}
			}
		}
		rv = append(rv, A[i])
		ri = append(ri, i)
	}

	return maxInArray(max) % 1000000007


}

func maxInArray(A []int) int {
	max := 0
	for _, v := range A{
		if v > max {
			max = v
		}
	}

	return max
}