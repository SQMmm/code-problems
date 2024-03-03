package gridchallenge

import "unicode/utf8"

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/grid-challenge/problem
 */

func gridChallenge(grid []string) string {
	if len(grid) == 0 {
		return "NO"
	}

	data := make([]string, len(grid[0]))
	for _, row := range grid {
		if len(row) != len(data) {
			return "NO"
		}
		sorted := sort(row)
		if !inOrder(data, sorted) {
			return "NO"
		}
		data = sorted
	}

	return "YES"
}

func inOrder(d, s []string)bool {
	for i := range s {
		if s[i] < d[i]{
			return false
		}
	}

	return true
}


func sort(row string) [] string {
	strs := make([]string, 0, len(row))
	for i, w := 0,0; i < len(row); i+=w {
		r, wight := utf8.DecodeRuneInString(row[i:])
		strs = append(strs, string(r))
		w = wight
	}

	return quickSort(strs)
}

func quickSort(s []string)[]string {
	if len(s) <=1 {
		return s
	}
	val := s[0]
	l, m, r := make([]string, 0),make([]string, 0),make([]string, 0)

	for _, ch := range s {
		if ch == val {
			m = append(m, ch)
			continue
		}
		if ch > val {
			l = append(l, ch)
			continue
		}
		r = append(r, ch)
	}

	return append(quickSort(l), append(m, quickSort(r)...)...)
}
