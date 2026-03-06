package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Create a test slice with various scenarios
	testWords := []string{
		"This", "is", "a", "apple",    // Case 1: Standard lowercase
		"and", "A", "orange",          // Case 2: Standard uppercase
		"it", "takes", "a", "hour",    // Case 3: The 'h' rule
		"but", "a", "banana",          // Case 4: Consonant (should NOT change)
		"is", "just", "a",             // Case 5: 'a' at the very end (boundary check)
	}

	fmt.Println("Before:", testWords)

	// 2. Call your function
	result := fixArticles(testWords)

	// 3. Print the result
	fmt.Println("After: ", result)
}

// fixArticles iterates through a slice and converts a/A to an/An 
// if the next word starts with a vowel or 'h'.
func fixArticles(words []string) []string {
	for i := 0; i < len(words); i++ {
		// Guard Clause: Skip if it's not our target word
		if words[i] != "a" && words[i] != "A" {
			continue
		}

		// Check if there is a next word and it isn't empty
		if i+1 < len(words) && len(words[i+1]) > 0 {
			
			// Get first character of next word as lowercase string
			firstChar := strings.ToLower(string(words[i+1][0]))

			// Check if it's a vowel or 'h'
			if strings.ContainsAny(firstChar, "aeiouh") {
				if words[i] == "a" {
					words[i] = "an"
				} else {
					words[i] = "An"
				}
			}
		}
	}
	return words
}
// func fixArticles(words []string) []string {
//     for i := 0; i < len(words); i++ {
// 		if words[i] == "a" || words[i] == "A" {
// 			if (i + 1) < len(words) {
// 				if len(words[i + 1]) != 0 {
// 					firstLetter := strings.ToLower(string(words[i+1][0]))
// 							if strings.ContainsAny(firstLetter, "aeiouh") {
// 								if words[i] == "a" {
// 									words[i] = "an"
// 								} else if words[i] == "A" {
// 									words[i] = "An"
// 								}
// 							}
// 						}
// 				}
				
// 			}
// 		}
// 	}
// 	return words
// }
