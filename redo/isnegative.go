package main

import "github.com/01-edu/z01"

// Write a function that prints 'T' (true) on a single line if the int passed as parameter is negative,
// otherwise it prints 'F' (false).

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
