package chessboardgame

/*
 * Complexity: Easy
 * Link: https://www.hackerrank.com/challenges/a-chessboard-game-1/problem
 */


const(
	min = int32(1)
	max = int32(15)
)

func chessboardGame(x int32, y int32) string {
	if x - 2 >= min && y + 1 <= max {
		w1 := whoWin(2, x - 2, y + 1, false)
		if w1 == true {
			return "First"
		}
	}

	if x - 2 >= min && y-1 >= min {
		w2 := whoWin(2, x - 2, y - 1, false)
		if w2 == true {
			return "First"
		}
	}

	if x + 1 <= max && y - 2 >= min {
		w3 := whoWin(2, x + 1, y - 2, false)
		if w3 == true {
			return "First"
		}
	}

	if x - 1 >= min && y - 2 >= min {
		w4 := whoWin(2, x - 1, y - 2, false)
		if w4 == true {
			return "First"
		}
	}
	return "Second"
}

func whoWin(n, x, y int32, isFirst bool) bool {
	if x - 2 >= min && y + 1 <= max {
		w1 := whoWin(n+1,x - 2, y + 1, !isFirst)
		if w1 == isFirst {
			return isFirst
		}
	}
	if x - 2 >= min && y-1 >= min {
		w2 := whoWin(n+1,x - 2, y - 1, !isFirst)
		if w2 == isFirst {
			return isFirst
		}
	}
	if x + 1 <= max && y - 2 >= min {
		w3 := whoWin(n+1,x + 1, y - 2, !isFirst)
		if w3 == isFirst {
			return isFirst
		}
	}
	if x - 1 >= min && y - 2 >= min {
		w4 := whoWin(n+1,x - 1, y - 2, !isFirst)
		if w4 == isFirst {
			return isFirst
		}
	}
	return !isFirst
}