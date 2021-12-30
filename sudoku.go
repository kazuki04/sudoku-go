package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// 0: not entered
// 1-9: entered
type Board [9][9]int

func prettify(b Board) string {
	// buffer
	var buf bytes.Buffer
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buf.WriteString("+---+---+---+\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				// Insert a vertical bar every three
				buf.WriteString("|")
			}
			// Convert type of int into string.
			buf.WriteString(strconv.Itoa(b[i][j]))
		}
		buf.WriteString("|\n")
	}
	buf.WriteString("+---+---+---+\n")
	return buf.String()
}

// duplicated checks the number of appearances.
// If number of appearances exceeds two return false.
func duplicated(c [10]int) bool {
	// k is number value
	for k, v := range c {
		// no check
		if k == 0 {
			continue
		}
		// v is number of appearances
		if v >= 2 {
			return true
		}
	}
	return false
}

// verify checks the three paterrens that rows, lines and 3x3.
func verify(b Board) bool {
	// check rows
	for i := 0; i < 9; i++ {
		//number of appearances
		var c [10]int
		for j := 0; j < 9; j++ {
			c[b[i][j]]++
		}
		if duplicated(c) {
			return false
		}
	}

	// check lines
	for i := 0; i < 9; i++ {
		// number of appearances
		var c [10]int
		for j := 0; j < 9; j++ {
			c[b[j][i]]++
		}
		if duplicated(c) {
			return false
		}
	}

	// check 3x3
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			var c [10]int
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					c[b[row][col]]++
				}
			}
			if duplicated(c) {
				return false
			}
		}
	}
	return true
}

// solved is the function judge whether the sudoku question is solved.
// It is solved when verify return true and all squares are filled.
// for loop in this function checks whether all squares are filled.
func solved(b Board) bool {
	if !verify(b) {
		return false
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

// backtrack is the algorithm.
// the argument b is passed by reference because of recursive call.
// As a result, squares of the board will be updated every time recursive call.
func backtrack(b *Board) bool {
	if solved(*b) {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				// if the square is 0 try whether the number is appropriate.
				// c is candidate.
				for c := 9; c >= 1; c-- {
					b[i][j] = c
					if verify(*b) {
						// if the verify return true explore deeply
						if backtrack(b) {
							return true
						}
					}
					b[i][j] = 0
				}
				return false
			}
		}
	}
	return false
}

func main() {
	b := Board{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Printf("%+v\n", prettify(b))

}
