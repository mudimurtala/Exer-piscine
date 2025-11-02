package main

import "github.com/01-edu/z01"

// Instructions
// Write a function that prints an int passed in parameter. All possible values 
// of type int have to go through. You cannot convert to int64.
// And its output :
// $ go run .
// -1230123
// $

func PrintNbr(n int) {

	if n == 0 {
		z01.PrintRune('0')
	}

	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			z01.PrintRune('9')
			n = -223372036854775808
		}
		n = -n
	}

	var digits []int
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n = n / 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune(digits[i] + '0'))
	}
}

func main() {
	PrintNbr(-123)
	z01.PrintRune('\n')
	PrintNbr(0)
	z01.PrintRune('\n')
	PrintNbr(123)
	z01.PrintRune('\n')
	PrintNbr(-9223372036854775808)
	z01.PrintRune('\n')
}