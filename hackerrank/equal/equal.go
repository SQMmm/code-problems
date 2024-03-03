package equal

import "fmt"

func TestEqual() {
	fmt.Println(equal([]int32{2, 2, 3, 7}), "=", 2)
	fmt.Println(equal([]int32{10, 7, 12}), "=", 3)
	fmt.Println(equal([]int32{1, 1, 5}), "=", 2)
}

func equal(arr []int32) int32 {
	if len(arr) == 0 {
		return 0
	}
	avOps := []int32{1, 2, 5}
	maxV := arr[0]
	for _, val := range arr {
		if val > maxV {
			maxV = val
		}
	}
	costs := make(map[int32][]int32, maxV+5)
	for _, op := range avOps {
		for i := int32(1); i <= maxV+5; i++ {
			diff := i - op
			if diff < 0 {
				continue
			}
			if diff == 0 {
				costs[i] = []int32{op}
				continue
			}
			costs[i] = append([]int32{op}, costs[diff]...)
		}
	}
	//fmt.Println(costs)

	
	
	num := maxV
	for {
		data := make(map[int32][]int32)
		opsNeeded := make([]int32, 0)
		for _, val := range arr {
			diff := num - val
			if diff == 0 {
				continue
			}
			cost := costs[diff]
			//fmt.Println(num, val, diff, cost)
			for _, op := range cost {
				if _, ok := data[op]; !ok {
					data[op] = make([]int32, 1)
				}
				lastOpI := 0
				if len(data[op]) > 0 {
					lastOpI = len(data[op]) - 1
				}
				data[op][lastOpI]++
			}

			opsNeeded = addNewOps(opsNeeded, costs[diff])
		}
		//fmt.Println(num, data)
		ok := true
		for _, ops := range data {
			for _, c := range ops {
				if c != int32(len(arr)-1) {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}

		if ok {
			return int32(len(opsNeeded))
		}
		num++

	}
}

func countOps(d map[int32][]int32) int32{
	//return int32(len(d))
	//res := make(map[int32]int32)
	fmt.Println(d)
	c := int32(0)
	for _, counts := range d{
		for i, count := range counts {
			if i == 0 {
				c += count
				continue

			}
			if count > counts[i-1] {
				c++
			}
			//diff := count - res[op]
			//res[op] += diff
		}
	}
	//for _, count := range res {
	//	c+=count
	//}

	return c
}

// func was made assuming that we can add values to any number of people
func equalTmp(arr []int32) int32 {
	newArr := make([]int32, 0, len(arr))
	mAr := make(map[int32]struct{}, len(newArr))
	maxV := int32(0)
	for _, val := range arr {
		if _, ok := mAr[val]; ok {
			continue
		}
		if val > maxV || maxV == 0 {
			maxV = val
		}
		mAr[val] = struct{}{}
		newArr = append(newArr, val)
	}

	//todo reduce space taken
	ops := make([]map[int32][]int32, 3)
	avOps := []int32{1, 2, 5}

	for j, op := range avOps {
		ops[j] = make(map[int32][]int32)
		for i := int32(1); i <= (maxV + 5); i++ {
			diff := i - op
			if diff < 0 {
				ops[j][i] = ops[j-1][i]
				continue
			}
			if diff == 0 {
				ops[j][i] = []int32{op}
				continue
			}
			//otherwise
			ops[j][i] = append(ops[j][diff], op)
		}
	}
	costs := ops[2]
	//fmt.Println(costs)
	avOps = append([]int32{0}, avOps...)
	data := make([]map[int32][]int32, len(newArr))
	for i, val := range newArr {
		data[i] = make(map[int32][]int32)
		for _, op := range avOps {
			max := maxV + op
			prev := make([]int32, 0)
			if i > 0 {
				prev = data[i-1][max]
			}
			diff := max - val
			if diff == 0 {
				data[i][max] = prev
				continue
			}

			//fmt.Println(max, prev, diff, costs[diff])

			data[i][max] = addNewOps(prev, costs[diff])

			//fmt.Println(data[i][max])
			//fmt.Println()

		}
	}

	min := -1
	for _, operations := range data[len(data)-1] {
		if len(operations) < min || min == -1 {
			min = len(operations)
		}
	}

	return int32(min)
}

func addNewOps(d, n []int32) []int32 {
	m := make(map[int32]int, len(n))
	for _, val := range n {
		m[val] += 1
	}

	for _, val := range d {
		m[val] -= 1
	}

	for val, count := range m {
		for i := 1; i <= count; i++ {
			d = append(d, val)
		}
	}

	return d
}
