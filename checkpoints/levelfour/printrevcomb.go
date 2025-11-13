// ## printrevcomb

// ### Instructions

// Write a program that prints in descending order on a single line all unique 
// combinations of three different digits so that the first digit is greater 
// than the second and the second is greater than the third.

// These combinations are separated by a comma and a space.

// ### Usage

// Here is an **incomplete** output :

// ```console
// $ go run . | cat -e
// 987, 986, 985, 984, 983, 982, 981, 980, 976, ..., 310, 210$
// $
// ```

// `999` or `000` are not valid combinations because the digits are not different.

// `789` should not be shown because the first digit is not greater than the second.




package main

import "github.com/01-edu/z01"

func main() {
	for a := 9; a >= 0; a-- {       // first digit
		for b := 9; b >= 0; b-- {   // second digit
			for c := 9; c >= 0; c-- { // third digit
				if a > b && b > c {
					// print the 3 digits
					z01.PrintRune(rune(a + '0'))
					z01.PrintRune(rune(b + '0'))
					z01.PrintRune(rune(c + '0'))

					// if not the last combination, print ", "
					if !(a == 2 && b == 1 && c == 0) {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n') // end line
}
