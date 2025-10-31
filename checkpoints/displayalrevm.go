package main

import "github.com/01-edu/z01"

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
