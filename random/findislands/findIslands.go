package findislands

import "fmt"


func TestFinIslands() {
	fmt.Println(findIslands([][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}), "= 6")
}

func findIslands(data [][]int) int {
	count := 0
	mIsl := make(map[string]struct{})
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			val := data[i][j]
			if val == 0 {
				continue
			}
			if _, ok := mIsl[fmt.Sprintf("%v%v", i, j)]; ok {
				continue
			}
			count++
			mIsl = findIsland(data, i, j, mIsl)
		}
	}

	return count
}
func findIsland(data [][]int, i, j int, mIsl map[string]struct{}) map[string]struct{} {
	//going right
	if (j + 1) == len(data[i]) || data[i][j+1] == 1 {
		key := fmt.Sprintf("%v%v", i, j+1)
		if _, ok := mIsl[key]; !ok {
			mIsl[key] = struct{}{}
			if (j+1) < len(data[i]) {
				mIsl = findIsland(data, i, j+1, mIsl)
			}
		}
	}

	//going down
	if (i+1) == len(data) || data[i+1][j] == 1 {
		key := fmt.Sprintf("%v%v", i+1, j)
		if _, ok := mIsl[key]; !ok {
			mIsl[key] = struct{}{}
			if (i+1) < len(data) {
				mIsl = findIsland(data, i+1, j, mIsl)
			}
		}
	}

	return mIsl
}
