package main

import (
	"strconv"
)

// processMarkers looks for (hex) and (bin) markers in the words slice.
// When it finds them, it converts the previous word to decimal
// and removes the marker from the slice.
func processMarkers(words []string) []string {
	for i := 0; i < len(words); i++ {

		// If marker appears at the beginning, remove it safely
		if i == 0 && (words[i] == "(hex)" || words[i] == "(bin)") {
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		// Handle hexadecimal conversion
		if words[i] == "(hex)" {
			// Convert previous word from base 16 to decimal
			num, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
			}

			// Remove the marker
			words = append(words[:i], words[i+1:]...)
			i--
		}

		// Handle binary conversion
		if words[i] == "(bin)" {
			// Convert previous word from base 2 to decimal
			num, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
			}

			// Remove the marker
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return words
}
