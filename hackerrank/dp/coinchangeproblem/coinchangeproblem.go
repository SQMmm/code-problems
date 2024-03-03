package coinchangeproblem

import "fmt"

/*
 * Complexity: Medium
 * Link: https://www.hackerrank.com/challenges/coin-change/problem
 */

func TestGetWays() {
	//fmt.Println(getWays(10, []int64{2, 5, 3, 6}), "=", 5)
	//fmt.Println(getWays(3, []int64{8, 3, 1, 2}), "=", 3)
	//fmt.Println(getWays(4, []int64{1, 2, 3}), "=", 4)
	fmt.Println(getWays(245, []int64{16, 30 ,9 ,17, 40, 13, 42, 5 ,25 ,49 ,7 ,23, 1, 44, 4 ,11, 33, 12 ,27, 2 ,38, 24 ,28, 32, 14, 50}), "=", 64027917156)
}

func getWays(n int32, c []int64) int64 {
	// Write your code here
	data := make([]int64, n+1)
	for _, val := range c {
		for i := int64(1); i <= int64(n); i++ {
			diff := i - val
			if diff < 0 {
				continue
			}
			if diff == 0 {
				data[i] += 1
				continue
			}
			if data[diff] == 0 {
				continue
			}
			data[i] +=  data[diff]
		}
	}

	return data[n]
}