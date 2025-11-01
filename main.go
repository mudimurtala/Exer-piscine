package main

import (
	"fmt"
	"os"
)

const N = 9 // size of the Sudoku (9x9)

// err prints "Error" and exits with a non-zero status.
// Used whenever input is invalid or the puzzle doesn't have exactly one solution.
func err() { fmt.Println("Error"); os.Exit(1) }

// validInput checks that we have exactly N strings and each string has length N.
// Returns true only when the command-line input matches the expected 9x9 shape.
func validInput(a []string) bool {
	if len(a) != N { // must provide 9 rows
		return false
	}
	for _, s := range a {
		if len(s) != N { // each row must have 9 characters
			return false
		}
	}
	return true
}

// ok checks whether placing 'val' at grid[row][col] is legal according to Sudoku rules.
// It returns false if the value already exists in the same row, column or 3x3 box.
func ok(g *[N][N]int, row, col, val int) bool {
	// Check row and column for duplicates
	for i := 0; i < N; i++ {
		if g[row][i] == val || g[i][col] == val {
			return false
		}
	}
	// Determine the top-left corner of the 3x3 box
	br, bc := row/3*3, col/3*3
	// Check the 3x3 box for duplicates
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if g[br+r][bc+c] == val {
				return false
			}
		}
	}
	return true
}

// solve performs recursive backtracking to find solutions.
// - g: pointer to current grid being worked on
// - sol: pointer where the first found full solution is copied
// - count: pointer to number of solutions found so far
//
// The function stops exploring when more than one solution is found (we only accept puzzles
// with exactly one solution). When a full valid grid is reached, it increments count and
// copies the grid into sol if it's the first solution.
func solve(g *[N][N]int, sol *[N][N]int, count *int) {
	// If more than one solution found already, no need to continue.
	if *count > 1 {
		return
	}

	// Find the first empty cell (scan row-major).
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if g[row][col] == 0 { // empty cell found
				// Try every possible value 1..9 in this cell
				for v := 1; v <= 9; v++ {
					if ok(g, row, col, v) { // if placing v is legal
						g[row][col] = v        // place it
						solve(g, sol, count)   // continue recursively
						g[row][col] = 0        // backtrack (undo)
					}
				}
				// After trying possibilities for this empty cell, return to backtrack
				// (we don't continue scanning other cells at this level)
				return
			}
		}
	}

	// No empty cells found => we reached a full valid solution.
	*count++ // increment number of solutions found
	// Save the first solution into sol
	if *count == 1 {
		*sol = *g
	}
}

func main() {
	// Read command-line arguments (skip program name)
	args := os.Args[1:]
	// Validate shape: exactly 9 arguments, each 9 chars long
	if !validInput(args) {
		err()
	}

	// grid holds the starting puzzle, sol will hold the first found solution
	var grid, sol [N][N]int

	// Parse input rows and populate grid.
	// Input characters: '.' means empty; '1'..'9' represent given digits.
	for row, s := range args {
		for col, ch := range s {
			if ch == '.' { // '.' means an empty cell; leave as 0
				continue
			}
			// Ensure character is a digit between '1' and '9'
			if ch < '1' || ch > '9' {
				err()
			}
			// Convert rune digit to integer value
			val := int(ch - '0')
			// Check that placing this given does not violate Sudoku rules
			if !ok(&grid, row, col, val) {
				err()
			}
			// Commit the given to the grid
			grid[row][col] = val
		}
	}

	// Find solutions using backtracking and count them.
	count := 0
	solve(&grid, &sol, &count)

	// We require exactly one solution. Otherwise print Error.
	if count != 1 {
		err()
	}

	// Print the solved grid: each row on its own line, numbers separated by single spaces.
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if col > 0 {
				fmt.Print(" ")
			}
			fmt.Print(sol[row][col])
		}
		fmt.Println()
	}
	// Extra newline at the end to match expected output style
	fmt.Println()
}
