package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	words := strings.Fields(string(content))
	words = processMarkers(words)
	words = handleUpTag(words)
	words = fixArticles(words)

	result := removeSpaceEfficient(strings.Join(words, " "))

	if err = os.WriteFile(outputFile, []byte(result), 0644); err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}

	fmt.Println("Processing complete! Output written to", outputFile)
}