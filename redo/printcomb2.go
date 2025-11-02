package main

import "github.com/01-edu/z01"

// Instructions
// Write a function that prints in ascending order and on a single line: 
// all possible combinations of two different two-digit numbers.
// These combinations are separated by a comma and a space.
// This is the incomplete output :
//
// $ go run . | cat -e
// 00 01, 00 02, 00 03, ..., 00 98, 00 99, 01 02, 01 03, ..., 97 98, 97 99, 98 99$
// $

func PrintComb2() {
	var l rune
	var e rune

	for l = 0; l <= 98; l++ {
		for e = l + 1; e <= 99; e++ {
			z01.PrintRune('0' + l/10)
			z01.PrintRune('0' + l%10)
			z01.PrintRune(' ')
			z01.PrintRune('0' + e/10)
			z01.PrintRune('0' + e%10)

			if (l != 98) || (e != 99) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintComb2()
}
