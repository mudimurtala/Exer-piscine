package main

import "strings"

func fixArticles(words []string) []string {

	vowels := "aeiouhAEIOUH"

	for i := 0; i < len(words)-1; i++ {

		if words[i] == "a" || words[i] == "A" {

			next := words[i+1]

			if len(next) > 0 && strings.ContainsRune(vowels, rune(next[0])) {

				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}

	return words
}