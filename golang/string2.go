package main

import "fmt"

func main() {
	text := "café"
	fmt.Println(text)
	fmt.Println(len(text))
	fmt.Println(text[0])
	fmt.Printf("%c\n", text[0])
	fmt.Printf("%c\n", text[1])

	for i := 0; i < len(text); i++ {
		fmt.Printf("%c ", text[i])
	}
	fmt.Println()

	for i, r := range text {
		fmt.Println(i, r)
	}
}
