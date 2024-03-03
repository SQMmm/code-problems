package goodlandelectricity

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * Complexity: Medium
 * Link: https://www.hackerrank.com/challenges/pylons/problem
 */

func TestPylons() {
	//fmt.Println(pylons(3, []int32{0, 1, 1, 1, 0, 0, 0}), "=", -1)
	//fmt.Println(pylons(3, []int32{0, 1, 0, 0, 0, 1, 1, 1, 1, 1}), "=", 3)
	//fmt.Println(pylons(2, []int32{0, 1, 0, 0, 0, 1, 0}), "=", -1)
	//fmt.Println(pylons(2, []int32{0, 1, 1, 1, 1, 0}), "=", 2)
	//fmt.Println(pylons(3, []int32{0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0}), "=", 4)
	strs := strings.Split(t, " ")
	arr := make([]int32, 0, 100000)
	for _, s := range strs {
		val, _ := strconv.Atoi(s)
		arr = append(arr, int32(val))
	}
	fmt.Println(pylons(4, arr), "=", 17901)

}

func pylons(k int32, arr []int32) int32 {
	// Write your code here
	maxDist := (k-1)*2
	z := make([][]int32, 0)
	zeros := int32(0)
	for _, val := range arr {
		if val == 0 {
			zeros++
			continue
		}
		if zeros > maxDist {
			return -1
		}
		// zeros == 1
		tmp := make([]int32, 2)
		tmp[0] = zeros
		z = append(z, tmp)
		if len(z) > 1 {
			z[len(z)-2][1] = zeros
		}
		zeros = 0
	}
	z[len(z)-1][1] = zeros

	if z[0][0] > (k-1) || z[len(z)-1][1] > (k-1) {
		return -1
	}

	i := 0
	sum := int32(0)
	for ; i < len(z); i++ {
		if i == 0 {
			sum += z[i][0]
		}else {
			sum += z[i][0] + 1
		}
		if sum > (k - 1) {
			break
		}
	}

	c := int32(1)
	i--
	for ; i < len(z); {
		kTmp := int32(0)
		j := i
		for ; j < len(z); j++ {
			if i == j {
				kTmp += z[j][1]
			}else {
				kTmp += z[j][1] + 1
			}
			if kTmp > maxDist {
				break
			}
		}
		c++
		i = j
		if j == (len(z) - 1) {
			return c
		}
	}

	return c
}
