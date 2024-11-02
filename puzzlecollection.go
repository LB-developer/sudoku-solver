package main

import (
	"encoding/json"
	"log"
	"os"
)

type Puzzles struct {
	Puzzles []Puzzle
}

type Puzzle struct {
	ID   string
	Grid [9][9]int
}

func LoadBoards() Puzzles {
	board, err := os.ReadFile("./puzzles.json")
	if err != nil {
		log.Fatalf("Couldn't read puzzles, err: %s", err)
	}

	var toSolve Puzzles
	if err = json.Unmarshal(board, &toSolve); err != nil {
		log.Fatalf("Couldn't unmarshal json of puzzle file: \n%s", err)
	}

	return toSolve
}
