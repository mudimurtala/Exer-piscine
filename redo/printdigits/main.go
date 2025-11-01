package main

import "github.com/01-edu/z01"

// Write a program that prints the decimal digits in ascending order (from 0 to 9) on a single line.
// A line is a sequence of characters preceding the end of line character ('\n').

func main() {
	for dig := '0'; dig <= '9'; dig++ {
		z01.PrintRune(dig)
	}
	z01.PrintRune('\n')
}
