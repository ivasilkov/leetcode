package _0036__Valid_Sudoku_

// 100% memory
func isValidSudoku(board [][]byte) bool {
	colSeen := [9][9]bool{}
	rowSeen := [9][9]bool{}
	sqrSeen := [3][3][9]bool{}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}
			v := board[row][col] - '0' - 1
			if colSeen[col][v] || rowSeen[row][v] || sqrSeen[col/3][row/3][v] {
				return false
			}
			colSeen[col][v] = true
			rowSeen[row][v] = true
			sqrSeen[col/3][row/3][v] = true
		}
	}
	return true
}

func isValidSudoku2(board [][]byte) bool {
	m := map[[3]byte]struct{}{}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			v := board[row][col]
			if v == '.' {
				continue
			}
			colKey := [3]byte{'c', byte(col), v}
			if _, ok := m[colKey]; ok {
				return false
			}
			rowKey := [3]byte{'r', byte(row), v}
			if _, ok := m[rowKey]; ok {
				return false
			}
			sqrKey := [3]byte{byte(row / 3), byte(col / 3), v}
			if _, ok := m[sqrKey]; ok {
				return false
			}
			m[colKey] = struct{}{}
			m[rowKey] = struct{}{}
			m[sqrKey] = struct{}{}
		}
	}
	return true
}

func isValidSudoku1(board [][]byte) bool {
	colMap := map[uint8]map[uint8]struct{}{}
	rowMap := map[uint8]map[uint8]struct{}{}
	sqrMap := map[uint8]map[uint8]struct{}{}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			v := board[row][col]
			if v == '.' {
				continue
			}
			if !addToMap(colMap, uint8(col), v) {
				return false
			}
			if !addToMap(rowMap, uint8(row), v) {
				return false
			}
			sqr := (row/3)*3 + (col / 3)
			if !addToMap(sqrMap, uint8(sqr), v) {
				return false
			}
		}
	}

	return true
}

func addToMap(m map[uint8]map[uint8]struct{}, k, v uint8) bool {
	if _, ok := m[k][v]; ok {
		return false
	}
	if _, ok := m[k]; !ok {
		m[k] = make(map[uint8]struct{})
	}
	m[k][v] = struct{}{}
	return true
}
