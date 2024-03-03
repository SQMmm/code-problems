package game2stacks

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * Complexity: Medium
 * Link: https://www.hackerrank.com/challenges/game-of-two-stacks/problem
 */

func twoStacks(maxSum int32, a []int32, b []int32) int32 {
	var (
		c1, c2, sum int32
	)

	for _, val := range a {
		if sum + val > maxSum {
			break
		}
		sum += val
		c1++
	}


	res := c1
	for _, num := range b {
		sum += num
		c2++
		for sum > maxSum && c1 > 0 {
			sum -= a[c1 - 1]
			c1--
		}

		if sum > maxSum {
			break
		}
		res = max(c1 + c2, res)
	}

	return res
}
func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func TestTwoStacks() {
	strs := strings.Split(t, "\n")
	//n, _ := strconv.Atoi(strs[0])
	strs =  strs[1:]

	strs = strs[15:18]
	fmt.Println(strs[0])
	fmt.Println(strs[1])
	fmt.Println(strs[2])
	//for i := 0; i < 3; i++ {
	//for i := 0; i < n; i++ {
		ss := strings.Split(strs[0], " ")
		maxSum, _ := strconv.Atoi(ss[2])

		a, b := make([]int32, 0), make([]int32, 0)
		ss = strings.Split(strs[1], " ")
		for _, s := range ss {
			num, _ := strconv.Atoi(s)
			a = append(a, int32(num))
		}

		ss = strings.Split(strs[2], " ")
		for _, s := range ss {
			num, _ := strconv.Atoi(s)
			b = append(b, int32(num))
		}

		fmt.Println(twoStacks(int32(maxSum), a, b))

		//if len(strs) <= 1 {
		//	break
		//}
		strs = strs[3:]
	//}
}
