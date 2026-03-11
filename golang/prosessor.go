package main

import (
	"fmt"
	"strconv"
	"strings"
)

func processMarkers(words []string) []string {
	for i := 0; i < len(words); i++ {
		// Skip if marker is at the beginning
		if i == 0 && (words[i] == "(hex)" || words[i] == "(bin)") {
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		} (hex) Simply add 42 (hex) and 10 (bin)
		
		if words[i] == "(hex)" {
			// Convert previous word from hex to decimal
			if num, err := strconv.ParseInt(words[i-1], 16, 64); err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
			}
			// Remove the marker
			words = append(words[:i], words[i+1:]...)
			i-- // Adjust index after removal
		} else if words[i] == "(bin)" {
			// Convert previous word from binary to decimal
			if num, err := strconv.ParseInt(words[i-1], 2, 64); err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
			}
			// Remove the marker
			words = append(words[:i], words[i+1:]...)
			i-- // Adjust index after removal
		}
	}
	return words
}

func main() {
	// Test Exercise 6
	words1 := []string{"10", "(bin)"}
	result1 := processMarkers(words1)
	fmt.Printf("Exercise 6 result: %v\n", result1) // Should print: [2]
	
	// Test Exercise 7
	words2 := []string{"42", "(hex)"}
	result2 := processMarkers(words2)
	fmt.Printf("Exercise 7 result: %v\n", result2) // Should print: [66]
	
	// Test Final Project
	text := "Simply add 42 (hex) and 10 (bin)"
	words := strings.Fields(text)
	result := processMarkers(words)
	finalText := strings.Join(result, " ")
	fmt.Printf("Final result: %s\n", finalText) // Should print: Simply add 66 and 2
}
