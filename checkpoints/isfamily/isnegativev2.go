package main

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
		z01.PrintRune('r')
		z01.PrintRune('u')
		z01.PrintRune('e')
	} else {
		z01.PrintRune('F')
		z01.PrintRune('a')
		z01.PrintRune('l')
		z01.PrintRune('s')
		z01.PrintRune('e')
	}
	z01.PrintRune('\n')
}

func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}
