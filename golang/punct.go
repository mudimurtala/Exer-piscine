package main

import (
	"fmt"
	"strings"
)

// AttachPunctuation attaches punctuation marks to the previous word in a slice of strings.
func AttachPunctuation(data []string) []string {
	punct := ".,!?;:"
	i := 0
	for i < len(data) {
		word := data[i]
		// punctuation should attach to previous word
		if strings.ContainsAny(word, punct) && i > 0 {
			data[i-1] = data[i-1] + word
			data = append(data[:i], data[i+1:]...)
			// Do not increment i, as the slice has shifted
		} else {
			i++
		}
	}
	return data
}

func main() {
	// Example test case
	input := []string{"Hello", ",", "world", "!", "How", "are", "you", "?"}
	result := AttachPunctuation(input)
	fmt.Println(result) // Output: [Hello, world! How are you?]
}
