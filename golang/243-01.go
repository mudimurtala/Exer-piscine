package main

import (
	"fmt"
	"strings"
)

func uppercaseLastN(words []string, n int) []string {
	if n <= 0 || len(words) == 0 {
		return words
	}

	if n > len(words) {
		n = len(words)
	}

	var start int
	start = len(words) - n
	for i := start; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}

	return words
}

func main() {
	words := []string{"this", "is", "so", "amazing"}
	fmt.Printf("uppercaseLastN(..., 2) -> %q\n", uppercaseLastN(words, 3))
} 