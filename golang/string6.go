package main

import (
    "fmt"
    "strings"
)

func main() {
	
}

func removeSpace (s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' && i+1 < len(runes) {
			if isPunctuation(runes[i+1]) {
				continue
			}
		}
	}
}
