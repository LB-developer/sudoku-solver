package main

func solve(row int, col int, board *[9][9]int, rows *MapOfSets, cols *MapOfSets, sections *MapOfSets) (*[9][9]int, bool) {
	// incrementation
	if col == 9 {
		row++
		col = 0
	}

	// base case
	if row == 9 {
		if solved := checkBoard(board); solved {
			return board, true
		} else {
			return nil, false
		}
	}

	// calculate section
	section := (row / 3) + 3*(col/3)

	// pre-set board numbers cannot be changed so we skip those
	if board[row][col] == 0 {
		for n := 1; n <= 9; n++ {
			// check if the current number is in the current row/col/section
			_, nInRow := (*rows)[row][n]
			_, nInCol := (*cols)[col][n]
			_, nInSection := (*sections)[section][n]

			if nInRow || nInCol || nInSection {
				continue
			}

			// add the current number to the set
			(*rows)[row][n] = struct{}{}
			(*cols)[col][n] = struct{}{}
			(*sections)[section][n] = struct{}{}

			// amend the current board cell with the current number
			board[row][col] = n

			// try and solve with the current number
			if _, solved := solve(row, col+1, board, rows, cols, sections); !solved {
				backtrack(rows, row, cols, col, sections, section, n, board)
			} else {
				return board, true
			}

		}
	} else {
		if _, solved := solve(row, col+1, board, rows, cols, sections); solved {
			return board, true
		}
	}

	return nil, false
}

func backtrack(r *MapOfSets, cr int, c *MapOfSets, cc int, s *MapOfSets, cs int, n int, b *[9][9]int) {
	// remove current number from the sets
	delete((*r)[cr], n)
	delete((*c)[cc], n)
	delete((*s)[cs], n)

	// reset the current board cell
	b[cr][cc] = 0
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
