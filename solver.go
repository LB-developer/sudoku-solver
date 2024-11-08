package main

func solve(row int, col int, board *[9][9]int, rows *MapOfSets, cols *MapOfSets, sections *MapOfSets) (*[9][9]int, bool) {
	if col == 9 {
		row++
		col = 0
	}

	if row == 9 {
		if solved := checkBoard(board); solved {
			return board, true
		} else {
			return nil, false
		}
	}

	if board[row][col] == 0 {
		section := row/3 + 3*(col/3)

		for n := 1; n <= 9; n++ {
			_, rFound := (*rows)[row][n]
			_, cFound := (*cols)[col][n]
			_, sFound := (*sections)[section][n]

			if rFound || cFound || sFound {
				continue
			}

			board[row][col] = n

			(*rows)[row][n] = struct{}{}
			(*cols)[col][n] = struct{}{}
			(*sections)[section][n] = struct{}{}

			if _, solved := solve(row, col+1, board, rows, cols, sections); solved {
				return board, true
			}
			// delete everything from their sets
			// reset board cell to 0
			delete((*rows)[row], n)
			delete((*cols)[col], n)
			delete((*sections)[section], n)
			board[row][col] = 0

		}
	} else {
		if _, solved := solve(row, col+1, board, rows, cols, sections); solved {
			return board, true
		}
	}

	return board, false
}

func checkBoard(b *[9][9]int) bool {
	for _, row := range b {
		for _, col := range row {
			if col == 0 {
				return false
			}
		}
	}

	return true
}
