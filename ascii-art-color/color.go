package main

import (
	"fmt"
	"os"
	"strings"
)
func parseArgs() (colorName string, substring string, text string, err error) {
    args := os.Args[1:]

    // No arguments
    if len(args) == 0 {
        err = fmt.Errorf("usage error")
        return
    }

    // Check if first arg is a color flag
    if strings.HasPrefix(args[0], "--color=") {
        colorName = strings.TrimPrefix(args[0], "--color=")

        if len(args) == 2 {
            // --color=red "string" → color the whole string
            text = args[1]
            substring = text  // whole string is the substring
        } else if len(args) == 3 {
            // --color=red kit "string" → color only 'kit'
            substring = args[1]
            text = args[2]
        } else {
            err = fmt.Errorf("usage error")
        }
        return
    }

    // No flag — original behavior
    if len(args) == 1 {
        text = args[0]
        return
    }

    err = fmt.Errorf("usage error")
    return
}

func getColorCode(name string) string {
    colors := map[string]string{
        "black":   "\033[30m",
        "red":     "\033[31m",
        "green":   "\033[32m",
        "yellow":  "\033[33m",
        "blue":    "\033[34m",
        "magenta": "\033[35m",
        "cyan":    "\033[36m",
        "white":   "\033[37m",
        "orange":  "\033[38;5;208m", // Added orange (256-color mode)
    }
    code, exists := colors[name]
    if !exists {
        return ""  // unknown color, no code
    }
    return code
}

func coloredPositions(text string, substring string) map[int]bool {
    positions := make(map[int]bool)
    if substring == "" {
        return positions
    }

    // Walk through the string looking for every occurrence of substring
    for i := 0; i <= len(text)-len(substring); i++ {
        if text[i:i+len(substring)] == substring {
            // Mark all positions in this match
            for j := 0; j < len(substring); j++ {
                positions[i+j] = true
            }
        }
    }
    return positions
}


func asciiArtColor(input string, line []string, colored map[int]bool, colorCode string) string {
	reset := "\033[0m"
	parts := strings.Split(input, "\n")
	var result string

	// We need to track the position of each character in the ORIGINAL
	// input string (before splitting), because coloredPositions was built
	// from the full text. We use offset to track where we are.
	offset := 0

	for i, part := range parts {
		if part == "" {
			if i < len(parts)-1 {
				result += "\n"
			}
			// A newline counts as 1 character in the original string
			offset++
			continue
		}

		for row := 0; row < 8; row++ {
			pos := offset // reset pos to start of this part for each row
			for _, char := range part {
				charArt := getChar(char, line)[row]

				if colored[pos] && colorCode != "" {
					result += colorCode + charArt + reset
				} else {
					result += charArt
				}
				pos++ // move to next character position
			}
			result += "\n"
		}

		offset += len(part) // advance offset past this part's characters
		if i < len(parts)-1 {
			result += "\n"
			offset++ // the \n separator counts as 1 character
		}
	}
	return result
}

