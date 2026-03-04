package main

import (
    "fmt"
    "strings"
)

func main() {
	fmt.Println(manipulateSentence("Mudi is an awesome engineer"))
}

func manipulateSentence (s string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])
	}
	result := strings.Join(words, " ")
	return result
}
