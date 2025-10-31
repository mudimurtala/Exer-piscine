package main

import "github.com/01-edu/z01"

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
