package main

func main() {
}

func solveSudoku(board [][]byte) {
	var backtracking func() bool
	backtracking = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}

				for k := '1'; k <= '9'; k++ {
					if isValid(i, j, byte(k), board) {
						board[i][j] = byte(k)
						if backtracking() {
							return true
						}
						board[i][j] = '.'
					}
				}
				//9个数都搜索完了，还没填满棋盘，则返回false
				return false
			}
		}
		return true
	}
	backtracking()

}

func isValid(row, col int, k byte, board [][]byte) bool {
	// 检查列
	for i := 0; i < 9; i++ {
		if board[i][col] == k {
			return false
		}

	}

	//检查行
	for j := 0; j < 9; j++ {
		if board[row][j] == k {
			return false
		}

	}

	// 检查3*3方格
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}

	return true
}
