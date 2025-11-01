package main

import "github.com/01-edu/z01"

// Write a function that prints, in ascending order and on a single line: all unique combinations 
// of three different digits so that, the first digit is lower than the second, and the second is 
// lower than the third.
// These combinations are separated by a comma and a space.
// This is the incomplete output :
//
// $ go run . | cat -e
// 012, 013, 014, 015, 016, 017, 018, 019, 023, ..., 689, 789$
// $
// 000 or 999 are not valid combinations because the digits are not different.
// 987 should not be shown because the first digit is not less than the second.

func PrintComb() {
	for l := '0'; l <= '7'; l++ {
		for t := l + 1; t <= '8'; t++ {
			for e := t + 1; e <= '9'; e++ {
				z01.PrintRune(l)
				z01.PrintRune(t)
				z01.PrintRune(e)

				if (l != '7') || (t != '8') || (e != '9') {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintComb()
}
