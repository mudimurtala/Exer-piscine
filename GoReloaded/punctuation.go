package main

import (
	"strings"
)

func removeSpaceEfficient(text string) string {
	// Fix space before each punctuation mark (DRY: one loop instead of 6 lines)
	for _, p := range []string{",", ".", "!", "?", ":", ";"} {
		text = strings.ReplaceAll(text, " "+p, p)
	}

	// Ensure a space always follows a comma
	text = strings.ReplaceAll(text, ",", ", ")

	// Fix ellipsis spacing: " ... " → "... "
	text = strings.ReplaceAll(text, " ... ", "... ")

	// Clean up any double spaces introduced by the steps above
	text = strings.ReplaceAll(text, "  ", " ")

	return strings.TrimSpace(text)
}