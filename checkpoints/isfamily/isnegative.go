package main

import "github.com/01-edu/z01"

// ✅ What does this program do?
// It checks whether a given number is negative.
// If the number is less than 0 → it prints T
// Otherwise → it prints F

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}

func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}
