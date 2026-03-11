package main

import (
	"strconv"
	"strings"
)

func handleUpTag(words []string) []string {
	for i := 1; i < len(words); i++ {
		if !strings.HasPrefix(words[i], "(") {
			continue
		}

		tag, numWords, consumed := parseTag(words, i)
		if tag == "" {
			continue
		}

		// Map each tag to its transformation function
		transforms := map[string]func(string) string{
			"(up":  strings.ToUpper,
			"(low": strings.ToLower,
			"(cap": capitalize,
		}

		// Find the matching transform by checking which key the tag starts with
		var transform func(string) string
		for prefix, fn := range transforms {
			if strings.HasPrefix(tag, prefix) {
				transform = fn
				break
			}
		}
		if transform == nil {
			continue
		}

		// Apply transform to the N words before the tag
		start := i - numWords
		if start < 0 {
			start = 0
		}
		for j := start; j < i; j++ {
			words[j] = transform(words[j])
		}

		// Remove the tag word(s) from the slice
		words = append(words[:i], words[i+consumed:]...)
		i--
	}
	return words
}

// parseTag reads a tag at startIdx. Tags can be one word "(up)" or
// two words "(low," + "3)". Returns tag, words-to-affect, words-consumed.
func parseTag(words []string, startIdx int) (string, int, int) {
	if !strings.HasPrefix(words[startIdx], "(") {
		return "", 0, 0
	}

	// Single-word tag: "(up)" or "(low,3)"
	if strings.Contains(words[startIdx], ")") {
		tag := extractTag(words[startIdx])
		if tag == "" {
			return "", 0, 0
		}
		return tag, extractNumber(tag), 1
	}

	// Two-word tag: "(low," + "3)"
	if startIdx+1 >= len(words) {
		return "", 0, 0
	}
	combined := words[startIdx] + " " + words[startIdx+1]
	tag := extractTag(combined)
	if tag == "" || !strings.Contains(tag, ",") {
		return "", 0, 0
	}
	return tag, extractNumber(tag), 2
}

// extractTag pulls out the "(…)" portion from a string.
func extractTag(word string) string {
	start := strings.Index(word, "(")
	end := strings.Index(word[start:], ")")
	if start == -1 || end == -1 {
		return ""
	}
	return word[start : start+end+1]
}

// extractNumber gets the N from tags like "(up, 3)". Defaults to 1.
func extractNumber(tag string) int {
	parts := strings.Split(strings.Trim(tag, "()"), ",")
	if len(parts) < 2 {
		return 1
	}
	n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 1
	}
	return n
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}