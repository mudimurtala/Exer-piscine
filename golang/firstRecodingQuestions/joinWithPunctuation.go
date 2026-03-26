package main

import (
	"fmt"
)

// Example - Input: ["hello", ",", "world", "!"] -> Expected Output: "hello, world!"
func joinWithPunctuation(tokens []string) string {
	result := ""
	for _, token := range tokens {
		if isPunctuation(token) {
			result += token
		} else {
			if result != "" {
				result += " "
			}
			result += token
		}
	}
	return result
}

func main() {
  greetings := []string{"hello", ",", "world", "!"}
  fmt.Printf("joinWithPunctuation(...) -> %q\n", joinWithPunctuation(greetings))
  sentence := []string{"Mudi", "!", "How", "are", "you", "?"}
  fmt.Printf("joinWithPunctuation(...) -> %q\n", joinWithPunctuation(sentence))
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
