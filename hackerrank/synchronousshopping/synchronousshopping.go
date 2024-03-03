package synchronousshopping

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func TestShop() {
	fmt.Println(shop(5, 5, []string{"1 1", "1 2", "1 3", "1 4", "1 5"}, [][]int32 {
		{1, 2, 10},{1,3, 10}, {2, 4, 10}, {3, 5, 10}, {4, 5, 10},
	}), "=", 30)
	//fmt.Println(shop(5, 3, []string{"1 1", "2 1 2", "2 2 3", "1 2", "1 2"}, [][]int32 {
	//	{1, 2, 10},{1,3, 15}, {2, 4, 1}, {2, 3, 10}, {3, 5, 5},
	//}), "=", 20)
}

// k - amount of fish types
func shop(n int32, k int32, centers []string, roads [][]int32) int32 {
	cs := make(map[int32][]int32, len(centers))
	for i, c := range centers {
		strs := strings.Split(c, " ")
		for _, s := range strs {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("failed to convert string to int: %v", err)
			}
			cs[int32(i)] = append(cs[int32(i)], int32(n))
		}
	}

	graph := make(map[int32][]int32)
	dur := make(map[string]int32)
	template := "%v-%v"
	for _, r := range roads {
		graph[r[0]] = append(graph[r[0]], r[1])
		graph[r[1]] = append(graph[r[1]], r[0])
		dur[fmt.Sprintf(template, r[0], r[1])] = r[2]
		dur[fmt.Sprintf(template, r[1], r[0])] = r[2]
	}

	needed := getNeeded(k, cs)

	queue := make([]int32, 0)
	durration := int32(0)
	for i := 0; i <= len(queue); i++ {
		//todo: check ig len(cs[queue[i]]) == 0
		ft := cs[queue[i]][1:]
		for _, t := range ft {
			if _, ok := needed[t]; ok {

			}
		}
	}

	return 0
}

func getNeeded(k int32, cs map[int32][]int32) map[int32]struct{} {
	ss :=  make([]int32, 0, cs[1][0] + cs[k][0])
	ss = append(ss, cs[1]...)
	ss = append(ss, cs[k]...)

	res := make(map[int32]struct{}, k)

	for i := int32(1); i <= k; i++ {
		res[i] = struct{}{}
	}

	for _, s := range ss {
		delete(res, s)
	}

	return res
}
