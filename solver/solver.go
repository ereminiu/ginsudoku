package solver

var n = 9
var row []map[int]bool
var col []map[int]bool
var box [][]map[int]bool

func findEmpty(board *[][]int) (int, int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (*board)[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func isValid(x, y, val int) bool {
	return !(row[x][val] || col[y][val] || box[x/3][y/3][val])
}

func solve(board *[][]int) bool {
	x, y := findEmpty(board)

	if x == -1 || y == -1 {
		return true
	}

	for val := 1; val < 10; val++ {
		if isValid(x, y, val) {
			(*board)[x][y] = val
			row[x][val] = true
			col[y][val] = true
			box[x/3][y/3][val] = true

			if solve(board) {
				return true
			}

			(*board)[x][y] = 0
			row[x][val] = false
			col[y][val] = false
			box[x/3][y/3][val] = false
		}
	}

	return false
}

func prepare() {
	row, col = make([]map[int]bool, n), make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		row[i] = make(map[int]bool)
		col[i] = make(map[int]bool)
	}
	box = make([][]map[int]bool, n/3)
	for i := 0; i < n/3; i++ {
		box[i] = make([]map[int]bool, n/3)
		for j := 0; j < n/3; j++ {
			box[i][j] = make(map[int]bool)
		}
	}
}

func solveSudoku(board [][]int) {
	prepare()

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val := board[i][j]
			if val == 0 {
				continue
			}

			row[i][val] = true
			col[j][val] = true
			box[i/3][j/3][val] = true
		}
	}

	solve(&board)
}

func CheckSolvable(board *[][]int) bool {
	prepare()

	return solve(board)
}

func GetSolution(board [][]int) [][]int {
	solveSudoku(board)

	return board
}
