package main

import "github.com/01-edu/z01"

// Write a program that prints the Latin alphabet in lowercase on a single line.
// A line is a sequence of characters preceding the end of line character ('\n').

func main() {
	for alph := 'a'; alph <= 'z'; alph++ {
		z01.PrintRune(alph)
	}
	z01.PrintRune('\n')
}
