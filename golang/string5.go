package main

import (
    "fmt"
    "strings"
)

func main() {
	text := "!!!Hello World!!!"
	cleanText := strings.Trim(text, "!")
	fmt.Println(cleanText)

	sentence := "apple-orange-banana"
	fruits := strings.Split(sentence, "-")
	fmt.Println(fruits)

	phrase := "The cat sat on the cat mat"
	newPhrase := strings.Replace(phrase, "sat", "jump", 1)
	fmt.Println(newPhrase)

	text2 := "The cat sat on the cat mat with another cat."
	newText2 := strings.ReplaceAll(text2, "cat", "dog")
	fmt.Println(newText2)

	marker := "(up, 3)"
	trimmed := strings.Trim(marker, "()")
	fmt.Println(trimmed)
	parts := strings.Split(trimmed, ",")
	fmt.Println(parts)
	word := strings.TrimSpace(parts[0])
	fmt.Println(word)

}

