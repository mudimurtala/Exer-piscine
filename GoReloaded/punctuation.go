package main

import (
	"strings"
)

func removeSpaceEfficient(text string) string {
	// Fix space before each punctuation mark
	for _, p := range []string{",", ".", "!", "?", ":", ";"} {
		text = strings.ReplaceAll(text, " "+p, p)
	}

	// Ensure a space always follows a comma
	text = strings.ReplaceAll(text, ",", ", ")

	// Fix ellipsis spacing: " ... " → "... "
	text = strings.ReplaceAll(text, " ... ", "... ")

	// Fix single-quote pairs: ' word ' → 'word'
	text = fixQuotePairs(text)

	// Clean up any double spaces introduced by the steps above
	text = strings.ReplaceAll(text, "  ", " ")

	return strings.TrimSpace(text)
}

// fixQuotePairs finds every pair of ' marks and trims the spaces
// immediately inside them, e.g. ' awesome ' → 'awesome'
func fixQuotePairs(text string) string {
	parts := strings.Split(text, "'")

	// If there are fewer than 3 parts, there's no complete pair (need open + content + close)
	if len(parts) < 3 {
		return text
	}

	for i := 1; i < len(parts)-1; i += 2 {
		parts[i] = strings.TrimSpace(parts[i]) // trim spaces inside the quote pair
	}

	return strings.Join(parts, "'")
}
