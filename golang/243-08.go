package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Q6 - Write a function that fixes spacing inside single quotes
// Fix spacing inside single quotes.
// Requirement:
// Declare the exact variables: 'result', 'match', and 'inner'.
// Example 1 - Input: "' awesome '" -> Expected Output: "'awesome'"
// Example 2 - Input: "' hello world '" -> Expected Output: "'hello world'"


func fixSingleQuotes(text string) string {
	result := text
	var match string
	var inner string

	re := regexp.MustCompile(`'([^']*)'`)

	result = re.ReplaceAllStringFunc(result, func(m string) string {
		match = m

		inner = match[1 : len(match) - 1]

		inner = strings.TrimSpace(inner)

		match = "'" + inner + "'"

		return match
	})

	return result
}

func main() {
  fmt.Printf("fixSingleQuotes(...) -> %q\n", fixSingleQuotes("' awesome '"))
  fmt.Printf("fixSingleQuotes(...) -> %q\n", fixSingleQuotes("' hello world '"))
}