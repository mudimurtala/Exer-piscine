package main

import (
	"fmt"
)

// Example 1 - Input: "," -> Expected Output: true
// Example 2 - Input: "!" -> Expected Output: true
// Example 3 - Input: "x" -> Expected Output: false


func isPunctuation(s string) bool {
	punctuations := []string{".", ",", "!", "?", ":", ";"}

	for _, p := range punctuations {
		if s == p {
			return true
		}
	}
	
	return false
}

func main() {
  fmt.Printf("isPunctuation(\",\") -> %v\n", isPunctuation(","))
  fmt.Printf("isPunctuation(\",\") -> %v\n", isPunctuation("!"))
  fmt.Printf("isPunctuation(\"x\") -> %v\n", isPunctuation("x"))
}
