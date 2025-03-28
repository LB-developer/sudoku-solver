package main

import "testing"

var invalidBoard = [9][9]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
}

var validBoard = [9][9]int{
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

var tests = map[string]struct {
	valid bool
	grid  *[9][9]int
}{
	"invalid": {
		valid: false,
		grid:  &invalidBoard,
	},
	"valid": {
		valid: true,
		grid:  &validBoard,
	},
}

func TestValidateBoard(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			expected := test.valid
			actual, _, _, _ := validator(test.grid)

			if expected != actual {
				t.Errorf("expected %v to be %v", actual, expected)
			}
		})
	}
}
