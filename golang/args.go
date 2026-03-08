package main

import (
	"fmt"
	"os"
	"strings" // Added this!
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run args.go [input] [output]")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// 1. READ
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading:", err)
		os.Exit(1)
	}

	// 2. TRANSFORM
	// We convert bytes to string -> Uppercase it -> Convert back to bytes
	transformed := toTitleCase(string(content))

	byteData := []byte(transformed)

	// 3. WRITE
	// Note: We use '=' instead of ':=' because err was already declared above
	err = os.WriteFile(outputFile, byteData, 0644)
	if err != nil {
		fmt.Println("Error writing:", err) // Specific error message
		os.Exit(1)
	}

	fmt.Printf("Success! Content written to %s\n", outputFile)
}

func toTitleCase(s string) string {
    words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		if len(words[i]) == 0 {
		firstChar := strings.ToUpper(string(words[i][0]))
		restCharacters := strings.ToLower(words[i][1:])
		words[i] = firstChar + restCharacters
		}
	}
	return strings.Join(words, " ")
}
