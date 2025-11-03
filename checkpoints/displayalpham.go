package main

import "github.com/01-edu/z01"

// ## displayalpham
// ### Instructions
// Write a program that displays the alphabet, with even letters in uppercase, and odd letters in lowercase, followed by a newline (`'\n'`).
//### Usage
// ```console
// $ go run . | cat -e
// aBcDeFgHiJkLmNoPqRsTuVwXyZ$
// $

// ✅ What does the program do as a whole?
// It loops from a to z
// And prints letters like:
// 📌 Odd positions → lowercase
// 📌 Even positions → uppercase

func main() {
	char := 0

	for i := 'a'; i <= 'z'; i++ {
		char++
		if char%2 == 0 {
			z01.PrintRune(i - 32)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
