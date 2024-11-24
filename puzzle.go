package main

type Puzzle struct {
	ID   string
	Grid [9][9]int
}

type MapOfSets map[int]map[int]struct{}

func validator(board *[9][9]int) (bool, *MapOfSets, *MapOfSets, *MapOfSets) {
	rows := make(map[int]map[int]struct{})
	cols := make(map[int]map[int]struct{})
	sections := make(map[int]map[int]struct{})

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			cell := board[r][c]

			/*
				a section is mapped to a coordinate i.e., 0 is top left, 8 is bottom right
				this works because all numbers divided by 3 (the edge of each section)
				are:
				0 when below 3,
				1 when above 3 but below 6,
				2 when above 6

				when the column moves to or beyond 3 we see why we multiply by 3
				... + 3 * (col 3 / 3) == 1 -- section 1
				... + 3 * (col 4 / 3) == 1 -- section 1
				... + 3 * (col 6 / 3) == 2 -- section 2

				full e.g. (top left)     (row 0 / 3) + 3 * (col 0 / 3) == 0
				full e.g. (bottom right) (row 8 / 3) + 3 * (col 8 / 3) == 8
			*/
			
			section := (r / 3) + 3*(c/3)
			
			// create a new set if it doesn't exist
			if rows[r] == nil {
				rows[r] = make(map[int]struct{})
			}

			if cols[c] == nil {
				cols[c] = make(map[int]struct{})
			}

			if sections[section] == nil {
				sections[section] = make(map[int]struct{})
			}
			if cell == 0 {
				continue
			}

			

			// check the set for the current cell
			if _, found := rows[r][cell]; found {
				return false, nil, nil, nil
			}

			if _, found := cols[c][cell]; found {
				return false, nil, nil, nil
			}

			if _, found := sections[section][cell]; found {
				return false, nil, nil, nil
			}

			// add cell to current row, col, section set
			rows[r][cell] = struct{}{}
			cols[c][cell] = struct{}{}
			sections[section][cell] = struct{}{}
		}
	}

	return true, (*MapOfSets)(&rows), (*MapOfSets)(&cols), (*MapOfSets)(&sections)
}
