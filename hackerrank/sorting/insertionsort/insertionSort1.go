package insertionsort


import (
	"fmt"
	"strconv"
	"strings"
)

func TestInsertion1() {
	insertionSort1(5, []int32{1, 2, 4, 5, 3})
	fmt.Println()
	insertionSort1(5, []int32{2, 4, 6, 8, 3})
}

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/insertionsort1/problem
 */

type ar []int32
func (a ar) String() string {
	strs := make([]string, 0, len(a))
	for _, num := range a {
		strs = append(strs, strconv.Itoa(int(num)))
	}

	return fmt.Sprintf("%s", strings.Join(strs, " "))
}

func insertionSort1(n int32, arr []int32) {
	val := arr[n - 1]
	a := ar(arr)
	for i := n - 1; i >= 1; i-- {
		if a[i-1] > val {
			a[i] = arr[i-1]
			fmt.Println(a)
		}else {
			a[i] = val
			fmt.Println(a)
			return
		}
	}

	a[0] = val
	fmt.Println(a)
}
