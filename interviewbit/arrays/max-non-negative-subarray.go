package arrays

func Maxset(A []int ) []int {
	B := make([]int, 0)
	C := make([]int, 0)
	s := 0
	c := 0
	for _, v := range A {
		if v >= 0 {
			B = append(B, v)
		}else{
			if (sum(B) > s) || (sum(B) == s && len(B) > c) {
				s, c, C = sum(B), len(B), B
			}

			B = make([]int, 0)
		}
	}

	if sum(B) > s {
		s, c, C = sum(B), len(B), B
	}else if len(B) > c {
		s, c, C = sum(B), len(B), B
	}

	return C
}

func sum(A []int) int {
	res := 0
	for _, v := range A {
		res += v
	}

	return res
}