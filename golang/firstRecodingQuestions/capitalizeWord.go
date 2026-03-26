package main

import (
	"fmt"
	"strings"
)

// Example 1 - Input: "hELLO" -> Expected Output: "Hello" 
// Example 2 - Input: "WORLD" -> Expected Output: "World" 


func capitalizeWord(word string) string {
	if len(word) == 0 {
		return ""
	}

	// Step 1: convert entire word to lowercase
	transformed := strings.ToLower(word)

	// Step 2: capitalize first letter
	return strings.ToUpper(string(transformed[0])) + transformed[1:]
}


func main() {
	fmt.Printf("capitalizeWord(\"hELLO\") -> %q\n", capitalizeWord("hELLO")) 
	fmt.Printf("capitalizeWord(\"WORLD\") -> %q\n", capitalizeWord("WORLD"))
	fmt.Printf("capitalizeWord(\"mudi\") -> %q\n", capitalizeWord("mudi"))
	fmt.Printf("capitalizeWord(\"mUDI\") -> %q\n", capitalizeWord("mUDI"))
}