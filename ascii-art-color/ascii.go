package main

import ("strings")

func asciiArt(input string, line []string) string {
	parts := strings.Split(input, "\n")
	var result string

	for i, part := range parts {
		if part == "" {
			if i < len(parts)-1 {
				result = result + "\n"
			}
			continue
		}

		for row := 0; row < 8; row++ {
			for _, char := range part {
				charLine := getChar(char, line)
				result += charLine[row]
			}
			result += "\n"
		}

	}
	return result
}

func getChar(c rune, line []string) []string {
	startLine := (int(c)-32)*9 + 1
	if startLine < 0 || startLine+8 > len(line) {
		return make([]string, 8)
	}
	return line[startLine : startLine+8]
}
