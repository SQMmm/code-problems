package castleofgrid

import (
	"fmt"
	"strings"
)

/*
 * Complexity: Medium
 * Link: https://www.hackerrank.com/challenges/castle-on-the-grid/problem
 */

func TestMinimumMoves() {
	//fmt.Println(minimumMoves([]string{"...", ".X.", "..."}, 0, 0, 1, 2), "==", 2)
	//fmt.Println(minimumMoves([]string{".X.", ".X.", "..."}, 0, 0, 0, 2), "==", 3)
	//fmt.Println(minimumMoves([]string{"....", "..X.", "....", ".X.."}, 1, 0, 3, 3), "==", 3)

	//fmt.Println(minimumMoves(strings.Split(t1, "\n"), 1, 0, 45, 45), "==", 79)
	fmt.Println(minimumMoves(strings.Split(t2, "\n"), 0, 0, 99, 99), "==", 2)
}

type point struct {
	x int32
	y int32
}


func minimumMoves(grids []string, sx, sy, gx, gy int32) int {
	startPoint := point{sx, sy}
	endPoint := point{gx, gy}
	queue := make([]point, 0)
	moves, usedPoints := getPossibleMoves(grids, startPoint, map[point]struct{}{startPoint: {}})
	queue = append(queue, moves...)

	c := 1
	nextLevelQueue := make([]point, 0)
	tmpUsedPoints := make([]point, 0)
	for i := 0; i < len(queue); i++ {
		p := queue[i]
		if p == endPoint {
			return c
		}

		m, u := getPossibleMoves(grids, p, usedPoints)
		tmpUsedPoints = append(tmpUsedPoints, mapToSlice(u)...)
		nextLevelQueue = append(nextLevelQueue, m...)

		if i == len(queue) - 1 {
			queue = append(queue, nextLevelQueue...)
			usedPoints = addToMap(usedPoints, tmpUsedPoints)
			c++
			nextLevelQueue = make([]point, 0)
			tmpUsedPoints = make([]point, 0)
		}
	}

	return c
}

func mapToSlice(m map[point]struct{}) []point {
	res := make([]point, 0, len(m))

	for p := range m {
		res = append(res, p)
	}

	return res
}

func addToMap(m map[point]struct{}, points []point) map[point]struct{} {
	for _, p := range points {
		m[p] = struct{}{}
	}
	return m
}

func getPossibleMoves(grids []string, p point, usedPoints map[point]struct{}) ([]point, map[point]struct{}) {
	n := int32(len(grids))
	res := make([]point, 0)

	// down
	if p.x < n - 1 {
		tmp := make([]point, 0)
		for i := p.x + 1; i < n; i++ {
			pp := point{i, p.y}
			if grids[i][pp.y] != '.' {
				break
			}
			if _, ok := usedPoints[pp]; !ok  {
				tmp = append(tmp, pp)
				usedPoints[pp] = struct{}{}
			}
		}
		for i := len(tmp) - 1; i >= 0; i-- {
			res = append(res, tmp[i])
		}
	}
	// up
	if p.x > 0 {
		tmp := make([]point, 0)
		for i := p.x - 1; i >= 0; i-- {
			pp := point{i, p.y}
			if grids[i][pp.y] != '.' {
				break
			}
			if _, ok := usedPoints[pp]; !ok {
				tmp = append(tmp, pp)
				usedPoints[pp] = struct{}{}
			}
		}
		for i := len(tmp) - 1; i >= 0; i-- {
			res = append(res, tmp[i])
		}
	}

	// to the right
	if p.y < n - 1 {
		tmp := make([]point, 0)
		for j := p.y + 1; j < n; j++ {
			pp := point{p.x, j}
			if grids[pp.x][j] != '.' {
				break
			}
			if _, ok := usedPoints[pp]; !ok {
				tmp = append(tmp, pp)
				usedPoints[pp] = struct{}{}
			}
		}
		for i := len(tmp) - 1; i >= 0; i-- {
			res = append(res, tmp[i])
		}
	}
	// up
	if p.y > 0 {
		tmp := make([]point, 0)
		for j := p.y - 1; j >= 0; j-- {
			pp := point{p.x, j}
			if grids[pp.x][j] != '.' {
				break
			}
			if _, ok := usedPoints[pp]; !ok {
				tmp = append(tmp, pp)
				usedPoints[pp] = struct{}{}
			}
		}
		for i := len(tmp) - 1; i >= 0; i-- {
			res = append(res, tmp[i])
		}
	}

	return res, usedPoints
}