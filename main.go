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
		} else {
			fmt.Printf("%s has been validated and will now be solved \n", boards.Puzzles[i].ID)
			fmt.Println("\nUnsolved board:")
			for _, row := range boards.Puzzles[i].Grid {
				fmt.Println(row)
			}
			solve(0, 0, &boards.Puzzles[i].Grid, rows, cols, sections)

			fmt.Println("===================")
			fmt.Println("Solved board:")
			for _, row := range boards.Puzzles[i].Grid {
				fmt.Println(row)
			}
		}
	}
}
