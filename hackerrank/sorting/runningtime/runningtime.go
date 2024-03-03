package runningtime

import "fmt"

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/runningtime/problem
 */

func TestRunningTime() {
	fmt.Println(runningTime([]int32{2, 1, 3, 1, 2}), "==", 4)
}

func runningTime(arr []int32) int32 {
	fn := func(arr []int32, i int32) int32 {
		val := arr[i]
		var shifts int32

		for i = i-1; i >= 0; i-- {
			if val < arr[i] {
				arr[i+1] = arr[i]
				shifts++
			}else {
				arr[i+1] = val
				return shifts
			}
		}
		arr[0] = val

		return shifts
	}

	var shifts int32
	for i := 1; i < len(arr); i++ {
		shifts += fn(arr, int32(i))
	}

	return shifts

}
