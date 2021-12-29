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
