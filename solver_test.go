package main

import (
	"reflect"
	"testing"
)

var unsolvedBoard = [9][9]int{
	{1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 9, 0, 0, 0, 3, 0},
	{0, 2, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 8, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 7, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 8, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 4, 0, 9, 0},
}

var solvedBoard = [9][9]int{
	{1, 3, 4, 2, 5, 6, 8, 7, 9},
	{5, 6, 7, 9, 1, 8, 2, 3, 4},
	{8, 2, 9, 3, 4, 7, 1, 5, 6},
	{2, 1, 3, 8, 7, 9, 4, 6, 5},
	{4, 9, 5, 1, 6, 2, 7, 8, 3},
	{6, 7, 8, 4, 3, 5, 9, 1, 2},
	{9, 4, 1, 5, 8, 3, 6, 2, 7},
	{7, 5, 2, 6, 9, 1, 3, 4, 8},
	{3, 8, 6, 7, 2, 4, 5, 9, 1},
}

var solverTests = map[string]struct {
	solvable bool
	unsolved *[9][9]int
	solved   *[9][9]int
}{
	"test_solve": {
		solvable: true,
		unsolved: &unsolvedBoard,
		solved:   &solvedBoard,
	},
}

func TestSolveBoard(t *testing.T) {
	for name, test := range solverTests {
		t.Run(name, func(t *testing.T) {
			valid, rows, cols, sections := validator(&unsolvedBoard)
			if !valid {
				t.Errorf("board aint valid")
				return
			}
			_, solved := solve(0, 0, test.unsolved, rows, cols, sections)

			if solved != test.solvable {
				t.Errorf("solver failed unexpectedly")
			}

			equal := reflect.DeepEqual(test.solved, test.unsolved)
			if !equal {
				t.Errorf("expected %v to be %v", test.unsolved, test.solved)
			}
		})
	}
}
