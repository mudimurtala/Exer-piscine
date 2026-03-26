package main

import (
	"fmt"
	"strings"
)

// Example 1 - Input: "apple" -> Expected Output: "an"
// Example 2 - Input: "horse" -> Expected Output: "an"
// Example 3 - Input: "book" -> Expected Output: "a"
// Example 4 - Input: "honest" -> Expected Output: "an" (starts with silent h)


func aOrAn(nextWord string) string {
	var firstLetter string
	vowelsAndH := []string{"a", "e", "i", "o", "u", "h"}

	if len(nextWord) == 0 {
		return "a"
	}

	firstLetter = strings.ToLower(string(nextWord[0]))

	for _, v := range vowelsAndH {
		if firstLetter == v {
			return "an"
		}
	}

	return "a"
}

func main() {
  fmt.Printf("aOrAn(\"apple\") -> %q\n", aOrAn("apple"))
  fmt.Printf("aOrAn(\"horse\") -> %q\n", aOrAn("horse"))
  fmt.Printf("aOrAn(\"book\") -> %q\n", aOrAn("book"))
  fmt.Printf("aOrAn(\"honest\") -> %q\n", aOrAn("honest"))
}