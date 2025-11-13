package main

import (
	"github.com/01-edu/z01"
)

func LastWord2(s string) string {
	result := ""

	// Step 1: Start from the last index (rightmost character)
	i := len(s) - 1

	// Step 2: Skip all trailing spaces
	for i >= 0 && s[i] == ' ' {
		i--
	}

	// Step 3: Find the start index of the last word
	end := i
	for i >= 0 && s[i] != ' ' {
		i--
	}
	start := i + 1 // the first letter of the last word

	// Step 4: Print and build result
	for j := start; j <= end; j++ {
		z01.PrintRune(rune(s[j]))     // print character
		result += string(rune(s[j]))  // add character to result
	}

	// Step 5: Print newline and add it to result
	z01.PrintRune('\n')
	result += "\n"

	return result
}

func main() {
	LastWord2("this        ...       is sparta, then again, maybe    not")
	LastWord2(" lorem,ipsum ")
	LastWord2(" ")
}
