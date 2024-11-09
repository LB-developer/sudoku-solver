# Sudoku Solver

A simple command-line Sudoku solver built in Go. 

This project validates if a given 9x9 sudoku puzzle is valid and then uses backtracking to solve.

The program uses predefined puzzles in JSON format, which are stored in the file `puzzles.json`. Each puzzle is represented as a 9x9 grid with 0s representing empty cells. Puzzles can be added or changed so long as their format remains the same; example of the format in puzzles.json:

```json
    {
      "puzzles": [
        {
          "id": "puzzle1",
          "grid": [
            [5, 3, 0, 0, 7, 0, 0, 0, 0],
            [6, 0, 0, 1, 9, 5, 0, 0, 0],
            [0, 9, 8, 0, 0, 0, 0, 6, 0],
            [8, 0, 0, 0, 6, 0, 0, 0, 3],
            [4, 0, 0, 8, 0, 3, 0, 0, 1],
            [7, 0, 0, 0, 2, 0, 0, 0, 6],
            [0, 6, 0, 0, 0, 0, 2, 8, 0],
            [0, 0, 0, 4, 1, 9, 0, 0, 5],
            [0, 0, 0, 0, 8, 0, 0, 7, 9]
          ]
        }
      ]
    }
```
## Adding or modifying puzzles
- Open the puzzles.json file in your preferred editor.
- Add new puzzle entries within the "puzzles" array, or modify the existing grids.

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/LB-developer/sudoku-solver
   cd sudoku-solver
   ```

2. **Build & run the project**:
    ```bash
    go build -o sudoku-solver
    ./sudoku-solver
    ```



