package shortestresch

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const weight = 6
func TestBFS()  {
	fmt.Println(bfs(5, 3, [][]int32{{1, 2}, {1, 3}, {3, 4}}, 1), "= [6 6 12 -1]" )
	fmt.Println(bfs(4, 2, [][]int32{{1, 2}, {1, 3}}, 1), "= [6 6 -1]" )
	fmt.Println(bfs(3, 1, [][]int32{{2, 3}}, 2), "= [-1 6]" )
	fmt.Println(bfs(5, 4, [][]int32{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 3), "= [12 6 6 12]" )
	fmt.Println(bfs(5, 4, [][]int32{{1, 3}, {2, 1}, {3, 4}, {4, 5}}, 3), "= [12 6 6 12]" )
	fmt.Println(bfs(8, 7, [][]int32{{2, 1}, {8, 5}, {3, 4}, {4, 5}, {2, 3}, {2, 5}, {7, 6}}, 3), "= [12 6 6 12 -1 -1 18]" )

	rows := strings.Split(tt, "\n")

	n, m := int32(0),int32( 0)
	edges := make([][]int32, 0)
	c := 0
	for i, r := range rows {
		if i == 0 {
			continue
		}
		strs := strings.Split(r, " ")
		if len(strs) == 1 {
			s, err := strconv.Atoi(strs[0])
			if err != nil {
				log.Fatalln(err)
			}
			//if c == 2 {
			fmt.Println("start")
			fmt.Println(bfs(n, m, edges, int32(s)))
			fmt.Println("end")
			//}
			n, m = 0, 0
			edges = make([][]int32, 0)
			c ++
			continue
		}
		n1, err := strconv.Atoi(strs[0])
		if err != nil {
			log.Fatalln(err)
		}
		n2, err := strconv.Atoi(strs[1])
		if err != nil {
			log.Fatalln(err)
		}
		if n == 0 {
			n, m = int32(n1), int32(n2)
			continue
		}
		edges = append(edges, []int32{int32(n1), int32(n2)})
	}

}

/*
 * Complexity: Medium
 * Link: https://www.hackerrank.com/challenges/bfsshortreach/problem
 */

func bfs(n int32, _ int32, edges [][]int32, ss int32) []int32 {
	// Write your code here
	links := make(map[int32][]int32)
	for _, e := range edges {
		links[e[0]] = append(links[e[0]], e[1])
		links[e[1]] = append(links[e[1]], e[0])
	}

	routes := make(map[int32]int32)
	queue := make([]int32, 0, len(edges))
	newQ := make([]int32, 0)
	queue = append(queue, links[ss]...)
	count := int32(1)
	for i := 0; i < len(queue); i++  {
		if _, ok := routes[queue[i]]; ok {
			if i == len(queue) - 1 {
				queue = append(queue, newQ...)
				newQ = make([]int32, 0)
				count++
			}
			continue
		}
		newQ = append(newQ, links[queue[i]]...)
		routes[queue[i]] = count * weight
		if i == len(queue) - 1 {
			queue = append(queue, newQ...)
			newQ = make([]int32, 0)
			count++
		}
	}

	res := make([]int32, 0, len(routes))
	for i := int32(1); i <= n; i++ {
		if i == ss {
			continue
		}
		if r, ok := routes[i]; ok {
			res = append(res, r)
			continue
		}
		res = append(res, -1)
	}

	return res
}

