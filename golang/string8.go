package main

import (
	"fmt"
	"strings"
)

func fixArticles(words []string) []string {
    for i := 0; i < len(words); i++ {
		if i == 'a' || 'A' {
			if (i + 1) < len(words) {
				nextWord := words[i + 1]
				if len(nextWord) != 0 {
					if nextWord[0] == 'a' || == 'e' || == 'i' || 'o' || == 'u' || == 'h' {
						if i == 'a' {
							i = 'an'
						} else if i == 'an' {
							i = 'An'
						}
					}
				}
			}
		}
	}
	return words
}


// 4. Step-by-Step Pseudocode
// Start a loop that goes through every word in the words slice using index i.

// Check if the current word is exactly "a" or "A".

// The "Safety Check": Make sure there is actually a "next word" to look at. (Check if i + 1 is less than the total length of the slice).

// The "Vowel/H Check":

// Get the next word: nextWord := words[i+1].

// Check if nextWord is not empty.

// Check the first character of nextWord.

// Is it 'a', 'e', 'i', 'o', 'u' or 'h' (case-insensitive)?

// The Modification:

// If the check is true:

// If the current word was "a", change it to "an".

// If the current word was "A", change it to "An".

// Return the modified words slice.

// 5. Deep Dive: Modifying the Slice
// In Go, when you pass a slice to a function, you are passing a "pointer" to the underlying data. This means if you do words[i] = "an", you are actually changing the original list! You don't need to build a "new bucket" like we did with the runes exercise, unless you want to be extra safe.

// Wait, what about "h"?
// The task says: "if the next word begins with a vowel... or a h". This is a common rule in some English dialects or specific words (like "an hour"). Your code should treat 'h' just like 'a, e, i, o, u'.