package main

import (
    "fmt"
	"strings"
)

func main() {
    // We print the result of the function call
    fmt.Println(removeSpaceEfficient("Mudi   , where are you going ?"))
}

func removeSpaceEfficient(s string) string {
    // Logic goes here
	var builder strings.Builder
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' && i+1 < len(runes) && strings.ContainsAny(string(runes[i+1]), ".,!?:;") {
			continue
		}

		builder.WriteRune(runes[i])
	}

	return builder.String()
}

