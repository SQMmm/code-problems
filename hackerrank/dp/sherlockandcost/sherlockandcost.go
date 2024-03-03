package sherlockandcost

func TestCost() {}

/*
 * Complexity: Medium
 * Link: https://www.hackerrank.com/challenges/sherlock-and-cost/problem
 */

func cost(b []int32) int32 {
	m1, m2, p1, p2 := int32(0), int32(0), int32(0), int32(0)
	for i := 1; i < len(b); i++ {
		m1 = max(p1+abs(b[i]-b[i-1]), p2+abs(b[i] - 1))
		m2 = p1 + abs(b[i-1] - 1)

		p1, p2 = m1, m2
	}

	return max(m1, m2)
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

func abs(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}