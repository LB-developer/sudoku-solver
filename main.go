package main

import (
	"fmt"
	"log"
)

func main() {
	boards := LoadBoards()

	for i := 0; i < len(boards.Puzzles); i++ {
		valid, rows, cols, sections := validator(&boards.Puzzles[i].Grid)
		if !valid {
			log.Printf("%s is not valid and will be skipped", boards.Puzzles[i].ID)
			continue
		}
		fmt.Printf("%s has been validated and will now be solved \n", boards.Puzzles[i].ID)
		solved, _ := solve(0, 0, &boards.Puzzles[i].Grid, rows, cols, sections)

		for _, row := range solved {
			fmt.Println(row)
		}
	}
}
