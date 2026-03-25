package main

import (
	"fmt"
)

// Example - Input: ["hello", ",", "world", "!"] -> Expected Output: "hello, world!"
func joinWithPunctuation(tokens []string) string {
	result := ""
	for i := 0; i < len(tokens); i++ {
		if isPunctuation(tokens[i]) {
			result += tokens[i]
		} else {
			if result != "" {
				result += " "
			}
			result += tokens[i]
		}
	}
	return result
}

func main() {
  tokens := []string{"hello", ",", "world", "!"}
  fmt.Printf("joinWithPunctuation(...) -> %q\n", joinWithPunctuation(tokens))
}

func isPunctuation(s string) bool {
	punctuations := []string{".", ",", "!", "?", ":", ";"}
	for _, p := range punctuations {
		if s == p {
			return true
		}
	}
	return false
}
