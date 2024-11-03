package main

import "log"

func main() {
	boards := LoadBoards()

	for i := 0; i < len(boards.Puzzles); i++ {
		valid, rows, cols, sections := BoardValidator(&boards.Puzzles[i].Grid)
		if !valid {
			log.Printf("%s is not valid and will be skipped", boards.Puzzles[i].ID)
			continue
		}

		log.Printf("%s is valid and will be solved with the following rows, cols, sections", boards.Puzzles[i].ID)
		log.Printf("rows: %v", rows)
		log.Printf("cols: %v", cols)
		log.Printf("sections: %v", sections)

		// SolveBoard(rows, cols, sections)

	}
}
