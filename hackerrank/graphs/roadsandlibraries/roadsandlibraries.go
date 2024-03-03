package roadsandlibraries

import (
	"fmt"
)

/*
 * Complexity: Medium
 * Link: https://www.hackerrank.com/challenges/torque-and-development/problem
 */

func TestRoadsAndLibraries() {
	fmt.Println(roadsAndLibraries(8, 3, 2, [][]int32{{1, 7}, {1, 2}, {1, 3}, {2, 3}, {5, 6}, {6, 8}}), "=", 19)
	fmt.Println(roadsAndLibraries(3, 2, 1, [][]int32{{1, 2}, {3, 1},{2, 3} }), "=", 4)
	fmt.Println(roadsAndLibraries(6, 2, 5, [][]int32{{1, 3}, {3, 4}, {2, 4}, {1, 2}, {2, 3}, {5, 6}}), "=", 12)
	//strs := strings.Split(t, "\n")
	//c := 0
	//for i := 0; i < len(strs); {
	//	row := strings.Split(strs[i], " ")
	//	n, _ := strconv.Atoi(row[0])
	//	m, _ := strconv.Atoi(row[1])
	//	cl, _ := strconv.Atoi(row[2])
	//	cr, _ := strconv.Atoi(row[3])
	//
	//	arr := make([][]int32, 0, m)
	//	for j := i + 1; j < m + 1 ; j ++ {
	//		row := strings.Split(strs[j], " ")
	//		n1, _ := strconv.Atoi(row[0])
	//		n2, _ := strconv.Atoi(row[1])
	//		arr = append(arr, []int32{int32(n1), int32(n2)})
	//	}
	//
	//	fmt.Println(roadsAndLibraries(int32(n), int32(cl), int32(cr), arr))
	//	i += m+1
	//	c++
	//}
}

func roadsAndLibraries(n int32, cLib int32, cRoad int32, cities [][]int32) int64 {
	if cLib < cRoad {
		return int64(cLib) * int64(n)
	}

	graph := make(map[int32][]int32, n * 2)
	for _, c := range cities {
		graph[c[0]] = append(graph[c[0]], c[1])
		graph[c[1]] = append(graph[c[1]], c[0])
	}


	needed := make(map[int32]struct{})
	for i := int32(1); i <= n; i++ {
		needed[i] = struct{}{}
	}

	libs := int64(0)
	roads := int64(0)
	for len(needed) > 0 {
		firstCity := getNeededCity(needed)
		queue := make([]int32, 0)
		queue = append(queue, graph[firstCity]...)
		delete(graph, firstCity)
		delete(needed, firstCity)
		libs++
		for i := 0; i < len(queue); i++ {
			c := queue[i]
			if _, ok := needed[c]; !ok {
				continue
			}
			queue = append(queue, graph[c]...)
			roads++
			delete(needed, c)
			delete(graph, c)
		}
	}


	return libs * int64(cLib) + roads * int64(cRoad)
}


func getNeededCity(cities map[int32]struct{}) int32 {
	for c := range cities {
		return c
	}

	return -1
}

//func getNeeded(cities [][]int32)

//func roadsAndLibraries(n int32, cLib int32, cRoad int32, cities [][]int32) int64 {
//	if cLib <= cRoad {
//
//		return int64(cLib) * int64(n)
//	}
//	graph := make(map[int32][]int32, n * 2)
//	for _, c := range cities {
//		graph[c[0]] = append(graph[c[0]], c[1])
//		graph[c[1]] = append(graph[c[1]], c[0])
//	}
//
//	needed := make(map[int32]struct{}, n)
//	for i := int32(1); i <= n; i++ {
//		needed[i] = struct{}{}
//	}
//
//	libs := int64(0)
//	roads := int64(0)
//	for len(needed) > 0 {
//		firstCity := getNeededCity(needed)
//		delete(needed, firstCity)
//		queue := make([]int32, 0)
//		queue = append(queue, graph[firstCity]...)
//		libs++
//
//		for i := 0; i < len(queue); i++ {
//			c := queue[i]
//			if _, ok := needed[c]; !ok {
//				continue
//			}
//			queue = append(queue, graph[c]...)
//			delete(needed, c)
//			delete(graph, c)
//			roads++
//		}
//	}
//
//	return libs * int64(cLib) + roads * int64(cRoad)
//}
//
//func getNeededCity(cities map[int32]struct{}) int32 {
//	for c := range cities {
//		return c
//	}
//
//	return -1
//}
