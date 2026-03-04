package main

import (
	"fmt"
	"unicode"
)

func main() {
	result := capitalize("mUDI")
	fmt.Println(result)
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)

	runes[0] = unicode.ToUpper(runes[0])

	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}

	return string(runes)
    
}


// 4. Pseudocode: Step-by-Step
// Here is how you should structure your algorithm:

// Accept the input string s.

// Handle the Edge Case: If the string is empty, return it immediately (otherwise, trying to access index 0 will crash your program).

// Convert the string into a []rune.

// Transform the Leader:

// Apply unicode.ToUpper specifically to the rune at index 0.

// Transform the Followers:

// Start a loop beginning at index 1 (not 0!).

// Continue until you reach the end of the slice.

// For every rune in this loop, apply unicode.ToLower.

// Rebuild: Convert the modified []rune back into a string.

// Return the result.