package main

import "github.com/01-edu/z01"

// ## displayalrevm
// ### Instructions
// Write a program that displays the alphabet in reverse, with even letters in uppercase, 
// and odd letters in lowercase, followed by a newline (`'\n'`).
// ### Usage
// $ go run . | cat -e
// zYxWvUtSrQpOnMlKjIhGfEdCbA$
// $

// ✅ What does the program print?
// It prints the alphabet from Z down to a, alternating case:
// 1st: z → odd → lowercase
// 2nd: y → even → uppercase
// 3rd: x → odd → lowercase
// 4th: w → even → uppercase

func main() {
	char := 0

	for i := 'z'; i >= 'a'; i-- {
		char++

		if char % 2 == 0 {
			z01.PrintRune(i - 32)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
