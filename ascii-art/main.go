package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	colorName, substring, text, err := parseArgs()
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
		return
	}

	if text == "" {
		return
	}

	data, err2 := os.ReadFile("standard.txt")
	if err2 != nil {
		fmt.Println("Error reading file:", err2)
		return
	}

	content := strings.Split(string(data), "\n")

	// Convert literal \n to real newline
	text = strings.ReplaceAll(text, `\n`, "\n")
	if substring != "" {
		substring = strings.ReplaceAll(substring, `\n`, "\n")
	}

	colorCode := getColorCode(colorName)
	colored := coloredPositions(text, substring)

	fmt.Print(asciiArtColor(text, content, colored, colorCode))
}









// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("It must not less than 2 argument")
// 		return
// 	}

// 	data, err := os.ReadFile("standard.txt")
// 	if err != nil {
// 		fmt.Println("Error found:", err)
// 		return
// 	}

// 	content := strings.Split(string(data), "\n")

// 	input := strings.Join(os.Args[1:], " ")
// 	input = strings.ReplaceAll(input, `\n`, "\n")

// 	fmt.Print(asciiArt(input, content))

// }












// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	// ── Step 1: Get the user's input word from the command line ──────────────
// 	// os.Args[0] is the program name, os.Args[1] is what the user typed
// 	if len(os.Args) < 2 {
// 		fmt.Println("Usage: go run . \"your text here\"")
// 		return
// 	}
// 	input := os.Args[1]
// 	fmt.Println("=== INPUT ===")
// 	fmt.Printf("You typed: %q\n\n", input)

// 	// ── Step 2: Read the entire banner file as one big string ────────────────
// 	// os.ReadFile reads everything at once — simple and clean
// 	// We no longer pass the filename as the argument; it's hardcoded here
// 	rawBytes, err := os.ReadFile("standard.txt")
// 	if err != nil {
// 		fmt.Println("Error reading standard.txt:", err)
// 		return
// 	}
// 	content := string(rawBytes)

// 	// ── Step 3: Split the file into blocks — one block per character ─────────
// 	// The key insight: each character's 8 rows are separated from the next
// 	// character by a BLANK LINE, which is just "\n\n" in the raw file.
// 	// Splitting on "\n\n" gives us one chunk per character.
// 	blocks := strings.Split(content, "\n\n")

// 	fmt.Println("=== BANNER FILE LOADED ===")
// 	fmt.Printf("Total character blocks found: %d\n", len(blocks))
// 	fmt.Printf("(We expect 95: from space ASCII 32 to ~ ASCII 126)\n\n")

// 	// ── Step 4: Split each block into its 8 individual rows ─────────────────
// 	// Each block looks like: "row1\nrow2\nrow3\n...row8"
// 	// Splitting on "\n" gives us a slice of 8 strings.
// 	// We store all of this in a 2D slice: banner[characterIndex][rowNumber]
// 	banner := make([][]string, len(blocks))
// 	for i, block := range blocks {
// 		banner[i] = strings.Split(block, "\n")
// 	}

// 	// Print a sample so you can see what the data looks like
// 	fmt.Println("=== SAMPLE: what the letter H looks like in memory ===")
// 	// 'H' is ASCII 72. Index = 72 - 32 = 40
// 	hIndex := int('H') - 32
// 	fmt.Printf("H is at banner index %d\n", hIndex)
// 	for rowNum, row := range banner[hIndex] {
// 		fmt.Printf("  row %d: %q\n", rowNum, row)
// 	}
// 	fmt.Println()

// 	// ── Step 5: Handle \n in the input — split into lines ───────────────────
// 	// If the user types "Hello\nThere", we treat the \n as a real line break.
// 	// strings.Split on "\n" gives us ["Hello", "There"]
// 	lines := strings.Split(input, "\\n")

// 	fmt.Println("=== INPUT LINES (after splitting on \\n) ===")
// 	for i, line := range lines {
// 		fmt.Printf("  line %d: %q\n", i, line)
// 	}
// 	fmt.Println()

// 	// ── Step 6: Print the ASCII art ─────────────────────────────────────────
// 	fmt.Println("=== ASCII ART OUTPUT ===")
// 	for lineIndex, line := range lines {
// 		// For an empty line (from \n\n in input), just print a blank line
// 		if line == "" {
// 			fmt.Println()
// 			continue
// 		}

// 		// For each of the 8 rows of height...
// 		for row := 0; row < 8; row++ {
// 			// ...go through every character in this line of text...
// 			for _, ch := range line {
// 				index := int(ch) - 32 // convert character to banner index

// 				// Safety check: only handle printable ASCII (32-126)
// 				if index < 0 || index >= len(banner) {
// 					fmt.Print("?")
// 					continue
// 				}

// 				// ...and print that character's art for this row
// 				// Make sure the banner slice has enough rows
// 				if row < len(banner[index]) {
// 					fmt.Print(banner[index][row])
// 				}
// 			}
// 			fmt.Println() // newline after each row
// 		}

// 		// After each line of text (except the last), print a blank line
// 		// only if there are more lines coming
// 		_ = lineIndex // suppress unused variable warning
// 	}
// }




// Run it step by step — what you will see

// When you run `go run . "Hi"`, the debug prints will show you:
// ```
// === INPUT ===
// You typed: "Hi"

// === BANNER FILE LOADED ===
// Total character blocks found: 95
// (We expect 95: from space ASCII 32 to ~ ASCII 126)

// === SAMPLE: what the letter H looks like in memory ===
// H is at banner index 40
//   row 0: " _    _  "
//   row 1: "| |  | | "
//   row 2: "| |__| | "
//   row 3: "|  __  | "
//   row 4: "| |  | | "
//   row 5: "|_|  |_| "
//   row 6: "         "
//   row 7: "         "

// === INPUT LINES (after splitting on \n) ===
//   line 0: "Hi"

// === ASCII ART OUTPUT ===
//  _    _   _  
// | |  | | (_) 
// | |__| |  _  
// |  __  | | | 
// | |  | | |_| 
// |_|  |_|









// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Error found")
// 	}

// 	data, err := os.ReadFile("standard.txt")
// 	if err != nil {
// 		fmt.Println("Error found:", err)
// 		return
// 	}

// 	// fmt.Println(data)
// 	// fmt.Printf("%q\n", data)
// 	// fmt.Printf("%T\n", data)

// 	content := strings.Split(string(data), "\n")

// 	// fmt.Println(content)
// 	// fmt.Printf("%T\n", content)
// 	// fmt.Printf("%q\n", content)

// 	input := strings.Join(os.Args[1:], " ")
// 	// fmt.Printf("%T\n", input)

// 	lines := strings.Split(input, "\\n")

// 	for i, line := range lines {
// 		if line == "" {
//             fmt.Println()
//             continue
//         }

// 		for row := 0; row < 8; row++ {
// 			for _, char := range input {
// 				charLine := getChar(char, content)
// 				fmt.Print(charLine[row])
// 			}
// 			fmt.Println()
// 		}

// 		if i < len(lines)-1 {
//             fmt.Println()
//         }
// 	}	
// }

// func getChar(c rune, line []string) []string {
// 	startline := (int(c) - 32) * 9
// 	return line[startline : startline+8]
// }
